package cmd

import (
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gogdl",
		Short: "Download all files in a folder from Google Drive.",
	}

	return cmd
}
