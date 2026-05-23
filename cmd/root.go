package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stash",
	Short: "Stash is a cli tool to create writeups using commands and their outputs",
	Long:  "Stash is a cli tool for writing reports using terminal responses esp. for ctfs",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		redBold("Oops. An error while executing Stash:", err)
		os.Exit(1)
	}
}
