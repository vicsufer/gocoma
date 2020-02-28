package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gocoma/models"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gocoma",
	Short: "gocoma is a tool for managing cloud environments",
	Long: `This tool is a go version of the Ruby gem 'Tacoma' https://github.com/pantulis/tacoma

If it is the first time running gocoma, please run the command:
  gocoma init`,
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Search config in home directory
	viper.AddConfigPath(home)
	viper.SetConfigName(".gocoma")
	viper.SetConfigType("yml")
	viper.AutomaticEnv() // read in environment variables that match

	// If another command from init is run without .gocomo.yml, show message.
	if err := viper.ReadInConfig(); err != nil && os.Args[1] != "init" {
		rootCmd.Help()
		fmt.Printf("The file .gocoma.yml is not detected, please run:\n  gocoma init\n")
		os.Exit(1)
	} else {
		models.Conf = &models.Configuration{}
		err = viper.Unmarshal(&models.Conf)
		if err != nil {
			fmt.Printf("unable to decode into config struct, %v", err)
		}
	}


}
