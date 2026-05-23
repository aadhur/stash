package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var addtitle string

var stashlogCmd = &cobra.Command{
	Use:   "log",
	Short: "Record it in .md",
	Long:  "Record the command and the response into the .md file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		command := strings.Join(args, " ")

		_, err := log(command, addtitle)
		if err != nil {
			return
		}

		greenBold("Stash logged")
	},
}

func init() {
	stashlogCmd.Flags().SetInterspersed(false)
	stashlogCmd.Flags().StringVarP(
		&addtitle,
		"CommandTitle",
		"c",
		"",
		"Adds command title",
	)
	rootCmd.AddCommand(stashlogCmd)
}
