package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gocoma/models"
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
			os.Exit(1)
		}
		selectedEnv := args[0]

		if models.Conf.SetCurrentEnv(selectedEnv) {
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
