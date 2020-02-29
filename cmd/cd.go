package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vicsufer/gocoma/models"
	"github.com/vicsufer/gocoma/utils"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change to the environment directory",
	Long: `Change to the directory configured as "Repo"
in the ~/.gocoma.yml file`,
	Run: func(cmd *cobra.Command, args []string) {

		var selectedEnv string
		//Check inputs
		if len(args) == 0 {
			selectedEnv = models.Conf.CurrentEnv
		} else {
			selectedEnv = args[0]
		}

		environments := models.Conf.Environments

		if env, ok := environments[selectedEnv]; !ok {
			fmt.Println("Selected environment not available, available environments are:")
			lsCmd.Run(cmd, []string{})
		} else {
			utils.SwitchShell(env.Repo)
		}
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
