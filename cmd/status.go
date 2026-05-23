package cmd

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check active session",
	Run: func(cmd *cobra.Command, args []string) {
		filename, err := activeSession()
		if err != nil {
			return
		}
		greenBold("Active session: " + filename)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
