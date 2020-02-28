package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gocoma/models"
)


// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available environments",
	Long: `To add new environments update the .gocoma.yml file in your home directory`,
	Run: func(cmd *cobra.Command, args []string) {

		for env := range models.Conf.Environments {
			currentEnvFlag := "    "
			if models.Conf.CurrentEnv == env {
				currentEnvFlag = "(*) "
			}
			fmt.Printf("%s%s\n", currentEnvFlag, env)
		}
	},
}
func init() {
	rootCmd.AddCommand(lsCmd)
}
