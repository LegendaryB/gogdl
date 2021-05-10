package main

import (
	"flag"
	"log"
	"os"

	"github.com/LegendaryB/gogdl/internal/download"
	"github.com/LegendaryB/gogdl/internal/gdrive"
)

func main() {
	folder := flag.String("folder", "", "Id of shared folder on Google Drive.")
	outdir := flag.String("outdir", "outdir", "Output directory for the downloads.")

	flag.Parse()

	service, err := gdrive.New()

	if err != nil {
		os.Exit(1)
	}

	log.Fatal(download.DriveFolder(service, *folder, *outdir))
}
