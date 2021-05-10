package gdrive

import (
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
)

const MIMETYPE_FOLDER = "application/vnd.google-apps.folder"

type DriveFolder struct {
	Id    string
	Name  string
	Files []*drive.File
}

func GetDriveFolder(service *drive.Service, id string) (*DriveFolder, error) {
	fileCall := service.Files.Get(id).
		SupportsAllDrives(true).
		SupportsTeamDrives(true)

	file, err := fileCall.Do()

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to retrieve metadata for resource: %s.", id)
	}

	err = ensureIsFolder(file)

	if err != nil {
		return nil, err
	}

	files, err := getFiles(service, id)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to retrieve children of resource", id)
	}

	return &DriveFolder{
		Id:    file.Id,
		Name:  file.Name,
		Files: files,
	}, err
}

func ensureIsFolder(file *drive.File) error {
	if file.MimeType != MIMETYPE_FOLDER {
		return errors.New(fmt.Sprintf("Resource with id '%s' is not a folder!", file.Id))
	}

	return nil
}

func getFiles(service *drive.Service, parentId string) ([]*drive.File, error) {
	var children []*drive.File
	var nextPageToken string

	for {
		query := fmt.Sprintf("'%s' in parents and mimeType != '%s' and trashed=false", parentId, MIMETYPE_FOLDER)

		listCall := service.Files.List().
			PageSize(100).
			SupportsAllDrives(true).
			SupportsTeamDrives(true).
			IncludeItemsFromAllDrives(true).
			IncludeTeamDriveItems(true).
			Fields("nextPageToken, files(id, name)").
			Q(query)

		if nextPageToken != "" {
			listCall.PageToken(nextPageToken)
		}

		files, err := listCall.Do()

		if err != nil {
			return nil, err
		}

		children = append(children, files.Files...)
		nextPageToken = files.NextPageToken

		if nextPageToken == "" {
			break
		}
	}

	return children, nil
}
