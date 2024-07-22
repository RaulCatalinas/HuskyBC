package utils

import (
	"os"
)

func createFolder(name string) {
	os.Mkdir(name, 0750)

	WriteMessage("info", "Created folder "+name)
}

func createFile(name string) {
	os.NewFile(0750, name)

	WriteMessage("info", "Created file "+name)
}

func CheckinFolderOrFile(path string, folder bool) {
	_, err := os.Stat(path)

	if err != nil {

		var folderType string
		if folder {
			folderType = "folder"
		} else {
			folderType = "file"
		}

		WriteMessage("error", "When checking "+folderType+": "+path)

		if os.IsNotExist(err) {
			WriteMessage("error", "Not found "+path)

			if folder {
				createFolder(path)
			} else {
				createFile(path)
			}
		} else {
			os.Exit(0)
		}
	}
}
