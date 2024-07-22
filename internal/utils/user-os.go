package utils

import (
	"os"
	"strings"
)

func createFolder(name string) {
	os.Mkdir(name, 0750)

	message := "Created folder " + name

	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: message,
	})
}

func createFile(name string) {
	os.NewFile(0750, name)

	message := "Created file " + name

	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: message,
	})
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

		WriteMessage(WriteMessageProps{
			Type:    "error",
			Message: GetErrorMessage("CheckingFolderOrFile"),
		})

		if os.IsNotExist(err) {
			errorMessage := GetErrorMessage("NotFound")

			WriteMessage(WriteMessageProps{
				Type:    "error",
				Message: strings.Replace(errorMessage, "{fileName}", path, -1),
			})

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
