package cmd

import (
	"github.com/spf13/cobra"
)

var restart bool

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Create the base .md file",
	Long:  "Create the base .md file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := start(args[0], restart)

		if err != nil {
			redBold("Error:", err)
			return
		}
		if restart {
			greenBold("Stash resumed: " + file + " found")
		} else {
			greenBold("Stash initiated: " + file + " created")
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(
		&restart,
		"restart",
		"r",
		false,
		"Restart and write to an exisiting .md file",
	)
}
