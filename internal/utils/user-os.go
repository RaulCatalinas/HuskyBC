package utils

import (
	"os"
	"strings"

	"github.com/RaulCatalinas/HuskyBC/internal/enums"
)

func createFolder(name string) {
	os.Mkdir(name, 0750)

	message := "Created folder " + name

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: message,
	})
}

func createFile(name string) {
	os.NewFile(0750, name)

	message := "Created file " + name

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: message,
	})
}

func CheckinFolderOrFile(path string, isFolder bool) {
	_, err := os.Stat(path)

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: GetErrorMessage("CheckingFolderOrFile"),
		})

		if os.IsNotExist(err) {
			errorMessage := GetErrorMessage("NotFound")

			WriteMessage(WriteMessageProps{
				Type:    enums.MessageTypeError,
				Message: strings.Replace(errorMessage, "{fileName}", path, -1),
			})

			if isFolder {
				createFolder(path)
			} else {
				createFile(path)
			}
		} else {
			os.Exit(0)
		}
	}
}
