﻿﻿<h1 align="center">gogdl</h1><div align="center">

[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![GitHub license](https://img.shields.io/github/license/LegendaryB/gogdl.svg?longCache=true&style=flat-square)](https://github.com/LegendaryB/gogdl/blob/master/LICENSE.md)

<sub>Built with ❤︎ by Daniel Belz</sub>
</div><br>

Simple CLI application to download all files in a folder from Google Drive. Team drives are also supported. At the moment only files at the top level are downloaded. Subfolders are ignored.

## Requirements

A Google Cloud Platform project with the Drive API enabled ([Guide](https://developers.google.com/drive/api/v3/quickstart/go#step_1_turn_on_the)).

You need to place the downloaded .json file next to the application and name it 'credentials.json'.

## Usage

### Download a folder
Just replace 'driveFolderId' with the Google Drive folder id you want to download. The files will be placed in currentFolder/outdir.

`sudo ./gogdl -folder driveFolderId`

### Download a folder to a custom location
Same as above but the files will be placed in currentFolder/mycustomfolder.

`sudo ./gogdl -folder driveFolderId -outdir mycustomfolder`
