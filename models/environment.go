package models

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
	"gocoma/utils"
)

type Environment struct {
	Aws_identity_file string `mapstructure: "aws_identity_file"`
	Aws_access_key_id string `mapstructure: "aws_access_key_id"`
	Aws_secret_access_key string`mapstructure: "aws_secret_access_key"`
	Region string `mapstructure: "region"`
	Repo string `mapstructure: "repo"`
}

func (environment *Environment) Activate() {
	setupSshConfig(environment)
	setupS3Cfg(environment)
	setupCredentials(environment)
	setupBoto(environment)
	setupEnvs(environment)
}

func setupSshConfig(environment *Environment) {
	if _, err := os.Stat( environment.Aws_identity_file ); os.IsNotExist(err) {
		fmt.Printf("The key %s does not exists.\n", environment.Aws_identity_file)
		os.Exit(1)
	}
	exec.Command("ssh-add", environment.Aws_identity_file).Output()
}

func setupCredentials(environment *Environment) {
	content := fmt.Sprintf(`
[default]
aws_access_key_id = %s
aws_secret_access_key = %s
region = %s
`, environment.Aws_access_key_id, environment.Aws_secret_access_key, environment.Region)

	home, _ := homedir.Dir()

	utils.WriteToFile(home+"/.aws/credentials", content)
}

func setupBoto(environment *Environment) {
	content := fmt.Sprintf(`
[credentials]
aws_access_key_id = %s
aws_secret_access_key = %s
`, environment.Aws_access_key_id, environment.Aws_secret_access_key)

	home, _ := homedir.Dir()
	utils.WriteToFile(home+"/.boto", content)
}

func setupS3Cfg(environment *Environment) {
	content := fmt.Sprintf(`
[default]
access_key = %s
secret_key = %s
`, environment.Aws_access_key_id, environment.Aws_secret_access_key)

	home, _ := homedir.Dir()
	utils.WriteToFile(home+"/.s3cfg", content)
}

func setupEnvs(environment *Environment) {
	os.Setenv("AWS_SECRET_ACCESS_KEY", environment.Aws_secret_access_key)
	os.Setenv("AWS_ACCESS_KEY_ID", environment.Aws_access_key_id)
}

