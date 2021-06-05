package cmd

import "github.com/spf13/cobra"

func CreateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gogdl",
		Short: "Download all files in a folder from Google Drive.",
	}

	cmd.AddCommand(createDownloadCommand())
	cmd.AddCommand(createVerifyCommand())

	return cmd
}
