package untar

import (
	"app/models"
	"errors"
	allFiles "github.com/A1311981684/utils/files"
	"github.com/A1311981684/utils/tar"
	"path/filepath"
)

func UnTarUpdate() error {
	//Get the DOWNLOADED package
	files, err := allFiles.GetAllFileNamesUnder(models.UpdatePath)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("too many update packages or not exist any")
	}
	ext := filepath.Ext(models.UpdatePath+files[0])
	if ext != ".update" {
		return errors.New("no valid *.update file found")
	}

	//Un-tar -> a directory contains update(name is the same as this project) + checksum gob file + scripts
	err = tar.UnTar(models.UpdatePath+files[0], models.NewFilesDirectory)
	if err != nil {
		return err
	}

	return nil
}
