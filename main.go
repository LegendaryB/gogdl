package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/LegendaryB/gogdl/gdrive"
	"github.com/avast/retry-go"
	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
)

func downloadFile(service *drive.Service, file *drive.File, dir string) error {
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

func downloadDir(service *drive.Service, folder string, outdir string) error {
	driveFolder, err := gdrive.GetDriveFolder(service, folder)

	if err != nil {
		return err
	}

	directory := filepath.Join(outdir, driveFolder.Name)
	err = os.MkdirAll(directory, os.ModePerm)

	if err != nil {
		return errors.Wrapf(err, "Failed to create directory '%s'.", directory)
	}

	for _, file := range driveFolder.Files {
		err = downloadFile(service, file, directory)

		if err != nil {
			return err
		}

		fmt.Printf("Download finished: %s/%s\n", driveFolder.Name, file.Name)
	}

	return nil
}

func main() {
	folder := flag.String("folder", "", "Id of shared folder on Google Drive.")
	outdir := flag.String("outdir", "outdir", "Output directory for the downloads.")

	flag.Parse()

	service, err := getService()

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(downloadDir(service, *folder, *outdir))
}
