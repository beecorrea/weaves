package cmd

import (
	"github.com/beecorrea/weaves/outfits/hackrun"
	"github.com/spf13/cobra"
)

var HackRunCmd = &cobra.Command{
	Use:   "hackrun <weave>",
	Short: "TUI for running a Hack script for a Weave, built with BubbleTea.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weave := args[0]
		hackrun.Wear(weave)
	},
}

func init() {
	MirandaCmd.AddCommand(HackRunCmd)
}
