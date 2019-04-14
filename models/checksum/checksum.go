package checksum

import (
	"app/models"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	files2 "github.com/A1311981684/utils/files"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Checksums map[string]string

var checksumMap Checksums

func CheckMD5s() error {
	err := loadChecksum()
	if err != nil {
		return err
	}

	err = CheckChecksums()
	if err != nil {
		return err
	}
	return nil
}

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
	//Try to get the extracted new files, the only directory extracted should be the root directory: projectName
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

	//Check if update package project name match current project name
	prjName, err := files2.GetCurrentProjectName()
	if err != nil {
		return err
	}
	if projectName != prjName {
		return errors.New("update package does not match")
	}

	//files, err = ioutil.ReadDir(models.NewFilesDirectory + projectName)
	var rangeFunc  func(path string) error
	rangeFunc = func(path string) error {
		files, err = ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		for _, v := range files {
			if v.IsDir() {
				err = rangeFunc(path + string(filepath.Separator)+v.Name())
				if err != nil {
					return err
				}
			}else {
				//Get extracted checksum
				if md5Ext, ok := checksumMap[strings.Replace(path + string(filepath.Separator)+v.Name(), models.NewFilesDirectory,
					string(filepath.Separator), 1)];ok {
					log.Println("md5Ext is", md5Ext)
					err = CheckChecksum(path + string(filepath.Separator)+v.Name(), md5Ext)
					if err != nil {
						return err
					}
				}else {
					log.Println("find checksum in map failed:",strings.Replace(path + string(filepath.Separator)+v.Name(), models.NewFilesDirectory,
						string(filepath.Separator), 1))
				}

			}
		}
		return nil
	}

	return rangeFunc(models.NewFilesDirectory + projectName)
}

func CheckChecksum(filePath, checksum string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	datas, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//calculate the file MD5
	h := md5.New()
	bytesMD5 := h.Sum(datas)
	strMD5 := fmt.Sprintf("%x", bytesMD5)
	log.Println("Calculated MD5:", strMD5, "compare to this MD5:", checksum)

	if strMD5 != checksum {
		return errors.New("MD5 checksum not match")
	}
	return nil
}
