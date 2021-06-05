package download

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/LegendaryB/gogdl/internal/gdrive"
	"github.com/avast/retry-go"
	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
)

func DriveFolder(service *drive.Service, folder string, outdir string) error {
	driveFolder, err := gdrive.GetDriveFolder(service, folder)

	if err != nil {
		return err
	}

	fmt.Printf("Drive folder: %s\n", driveFolder.Name)

	directory := filepath.Join(outdir, driveFolder.Name)
	err = os.MkdirAll(directory, os.ModePerm)

	if err != nil {
		return errors.Wrapf(err, "Failed to create directory '%s'.", directory)
	}

	for _, file := range driveFolder.Files {
		err = downloadDriveFile(service, file, directory)

		if err != nil {
			return err
		}

		fmt.Printf("\t%s -> done\n", file.Name)
	}

	return nil
}

func downloadDriveFile(service *drive.Service, file *drive.File, dir string) error {
	return retry.Do(func() error {
		path := filepath.Join(dir, file.Name)

		response, err := service.Files.Get(file.Id).
			SupportsAllDrives(true).
			SupportsTeamDrives(true).
			Download()

		if err != nil {
			return errors.Wrapf(err, "Failed to download file with id '%s'.", file.Id)
		}

		f, err := os.Create(path)

		if err != nil {
			return errors.Wrapf(err, "Failed to create file '%s'.", file.Name)
		}

		defer f.Close()

		_, err = io.Copy(f, response.Body)

		if err != nil {
			return errors.Wrapf(err, "Failed to write file '%s'.", file.Name)
		}

		return nil
	}, retry.Attempts(3))
}
