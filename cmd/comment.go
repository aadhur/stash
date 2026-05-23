package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var status string

var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Note it down in .md",
	Long:  "Note down any points as commands or secondary titles with -t",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		command := strings.Join(args, " ")

		_, err := comment(command, status)

		if err != nil {
			redBold(err)
			return
		}

		greenBold("Comment Recorded")
	},
}

func init() {
	commentCmd.Flags().StringVarP(
		&status,
		"title",
		"t",
		"",
		"Adds secondary title",
	)
	rootCmd.AddCommand(commentCmd)
}
