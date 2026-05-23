package cmd

import (
	"github.com/spf13/cobra"
)

var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End the logging session",
	Long:  "Ending the session means you can't edit the .md file anymore",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		endmsg, err := end()
		if err != nil {
			return
		} else {
			greenBold(endmsg)
		}

	},
}

func init() {
	rootCmd.AddCommand(endCmd)
}
