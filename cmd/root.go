package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

const NodepingTokenKey = "NODEPING_TOKEN"
var nodepingToken string

var rootCmd = &cobra.Command{
	Use: "nodeping-cli",
	Short: "Script for making calls to nodeping.com",
	Long:  "Script for making calls to nodeping.com",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terraform-enterprise-migrator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	foundError := false

	// Get Nodeping Token from env vars
	nodepingToken = os.Getenv(NodepingTokenKey)

	if nodepingToken == "" {
		fmt.Println("Error: Environment variable for NODEPING_TOKEN is required to execute plan and migration")
		fmt.Println("")
		foundError = true
	}

	if foundError {
		os.Exit(1)
	}
}

