package main

import (
	"os"

	"github.com/LegendaryB/gogdl/cmd"
)

func main() {
	cmd := cmd.Root()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	/*folder := flag.String("folder", "", "Id of shared folder on Google Drive.")
	outdir := flag.String("outdir", "outdir", "Output directory for the downloads.")

	flag.Parse()

	service, err := gdrive.New()

	if err != nil {
		os.Exit(1)
	}

	log.Fatal(download.DriveFolder(service, *folder, *outdir)) */
}
