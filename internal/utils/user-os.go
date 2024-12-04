package utils

import (
	"os"
	"strings"

	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	errorMessages "github.com/RaulCatalinas/HuskyBC/internal/error_messages"
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
		Type:    enums.MessageTypeSuccess,
		Message: message,
	})
}

func CheckinFolderOrFile(path string, isFolder bool) {
	_, err := os.Stat(path)

	if err != nil {
		if isFolder {
			createFolder(path)
		} else {
			createFile(path)
		}

		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.CheckingFolderOrFileError],
		})

		if os.IsNotExist(err) {
			errorMessage := errorMessages.FILE_ERROR_MESSAGES[enums.NotFoundError]

			WriteMessage(WriteMessageProps{
				Type:    enums.MessageTypeError,
				Message: strings.Replace(errorMessage, "{fileName}", path, -1),
			})
		} else {
			os.Exit(0)
		}
	}
}
