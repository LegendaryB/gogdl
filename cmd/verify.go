package cmd

import (
	"fmt"

	"github.com/LegendaryB/gogdl/internal/download"
	"github.com/LegendaryB/gogdl/internal/gdrive"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func createVerifyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "download [folderId]",
		Short:        "Download files from a Google Drive folder",
		Long:         "Use this command to download all files from a Google Drive folder.",
		SilenceUsage: true,
		Args:         cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("folder id not provided but required")
			}

			folderId := args[0]
			output, _ := cmd.Flags().GetString("output")

			service, err := gdrive.NewService()

			if err != nil {
				return errors.Wrap(err, "failed to initialize google drive service")
			}

			download.DriveFolder(service, folderId, output)

			if err != nil {
				return errors.Wrap(err, "failed to download files")
			}

			fmt.Println("Done")

			return nil
		},
	}

	cmd.Flags().String("output", "./out", "Outasdasdsa")

	return cmd
}
