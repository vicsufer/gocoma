/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Select the environment to use",
	Long: `Switch between the different environments configured in the .gocoma.yml
file in the home directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Check inputs
		if len(args) < 1 {
			fmt.Println("No environment selected")
			os.Exit(0)
		}
		selectedEnv := args[0]

		if Conf.SetCurrentEnv(selectedEnv) {
			fmt.Printf("Environment changed to %s\n", Conf.CurrentEnv)
		} else {
			fmt.Println("Selected environment not available, available environments are:")
			lsCmd.Run(cmd, []string{})
		}

	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
