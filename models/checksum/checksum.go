package checksum

import (
	"app/models"
	"encoding/gob"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
)

type Checksums map[string]string

var checksumMap Checksums

func loadChecksum() error {
	f, err := os.Open(models.NewFilesDirectory + "CHECKSUM")
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	decoder := gob.NewDecoder(f)
	err = decoder.Decode(&checksumMap)
	if err != nil {
		return err
	}

	return nil
}

func CheckChecksums() error {
	if checksumMap == nil || len(checksumMap) == 0 {
		return errors.New("checksum map is empty or not loaded")
	}
	files, err := ioutil.ReadDir(models.NewFilesDirectory)
	if err != nil {
		return err
	}
	projectName := ""
	for _, v := range files {
		if v.IsDir() {
			projectName = v.Name()
			break
		}
	}
	if projectName == "" {
		return errors.New("no extracted update directory found")
	}

	files, err = ioutil.ReadDir(models.NewFilesDirectory + projectName)
	return nil
}
