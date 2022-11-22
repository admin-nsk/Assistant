package main

import (
	"github.com/admin-nsk/Assistant/note_service/internal/pkg/path"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var RootCmd = &cobra.Command{
	Use:   "note_service",
	Short: "note",
}

func init() {
	RootCmd.SuggestionsMinimumDistance = 1
	RootCmd.SilenceUsage = true
	RootCmd.PersistentFlags().StringP("config", "c", path.EtcDir(),
		"Path to configuration file or directory with config files")
}

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}
