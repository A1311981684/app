package models

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var BackupPath = ""
var UpdatePath = ""
var WorkingDirectory = ""
var NewFilesDirectory = ""
var Separator = ""

//Path initialization
func init(){

	//Determine the running operating system to assign the Separator
	switch runtime.GOOS {
	case "windows":
		Separator = "\\"
	case "linux":
		Separator = "/"
	default:
		log.Fatal("unsupported operating system: " + runtime.GOOS)
	}
	log.Println("Programmer is running on", runtime.GOOS)

	//Get working directory to set update or backup path
	WorkingDirectory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Working directory is:", WorkingDirectory)

	BackupPath = WorkingDirectory + Separator + "BACKUPS" + Separator
	UpdatePath = WorkingDirectory + Separator + "UPDATES" + Separator + "DOWNLOADS" + Separator
	NewFilesDirectory = WorkingDirectory + Separator + "UPDATES" + Separator + "FILES" + Separator

	log.Println("Backups directory is:", BackupPath)
	log.Println("Updates directory is:", UpdatePath)
	log.Println("Extracted files directory:", NewFilesDirectory)
}

/*
Sequences:
1 Un-tar the update package from DOWNLOADS to the FILES
2 Read the CHECKSUM gob file
3 For each file name in the gob, calculate and verify its MD5 and compare with the MD5 that included in the gob
4 If all the checksums are checked and matched, continue.Else, Return error
5 Copy all the files that are going to be replaced by new files to the BACKUPS directory
6 Apply new files to cover old files.If error occurred, Recover copied backup files back to their original location
 */
