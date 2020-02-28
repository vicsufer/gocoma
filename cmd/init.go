/*
Copyright © 2020 NAME HERE @vicsufer

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gocoma/models"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate the default gocoma.yml",
	Long: `This command will generate .gacoma.yml at your home directory.

The file .gacoma.yml contains a templated environment, update the values as you wish, 
you can also add as many environents as needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				//Create a mapping with an empty env to create the template file
				var emptyConfig = models.Configuration{
					CurrentEnv: "sample_environment",
					Environments: map[string]models.Environment{"sample_environment": models.Environment{}},
				}
				viper.Set("currentenv", emptyConfig.CurrentEnv)
				viper.Set("environments", emptyConfig.Environments)

				//Safe write to avoid destroying already configured file
				viper.SafeWriteConfig()
				fmt.Println("The .gocoma.yml file has been created at your home directory!")
			}
		} else {
			fmt.Println("The file .gocoma.yml already exists at your home directory")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
