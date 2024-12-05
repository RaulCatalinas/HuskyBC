package utils

import (
	"os"
	"strings"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	errorMessages "github.com/RaulCatalinas/HuskyBC/internal/error_messages"
)

func modifyNpmIgnore(filesToAdd interface{}) {
	defer func() {
		if r := recover(); r != nil {
			WriteMessage(WriteMessageProps{
				Type:    enums.MessageTypeError,
				Message: errorMessages.FILE_ERROR_MESSAGES[enums.NpmIgnoreWriteError],
			})

			os.Exit(1)
		}
	}()

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: "Writing in the file \".npmignore\"...",
	})

	CreateFolderOrFileIfNotExists(constants.PATH_DIR_NPMIGNORE, false)

	data := readFile(constants.PATH_DIR_NPMIGNORE)

	ignoredFiles := string(data)

	if ignoredFiles == "" {
		var filesToWrite string

		switch v := filesToAdd.(type) {
		case string:
			filesToWrite = v
		case []string:
			filesToWrite = strings.Join(v, "\n")
		default:
			WriteMessage(WriteMessageProps{
				Type:    enums.MessageTypeError,
				Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.InvalidTypeForFilesToAddError],
			})

			os.Exit(1)
		}

		writeFile(constants.PATH_DIR_NPMIGNORE, []byte(filesToWrite))

		WriteMessage(WriteMessageProps{
			Type:    "success",
			Message: "\".npmignore\" file modified successfully",
		})

		return
	}

	ignoredFilesArray := strings.FieldsFunc(ignoredFiles, func(r rune) bool {
		return r == '\n'
	})

	trimmedIgnoredFilesArray := make([]string, 0, len(ignoredFilesArray))

	for _, file := range ignoredFilesArray {
		trimmedFile := strings.TrimSpace(file)

		if trimmedFile != "" {
			trimmedIgnoredFilesArray = append(trimmedIgnoredFilesArray, trimmedFile)
		}
	}

	switch v := filesToAdd.(type) {
	case string:
		trimmedIgnoredFilesArray = append(trimmedIgnoredFilesArray, v)
	case []string:
		trimmedIgnoredFilesArray = append(trimmedIgnoredFilesArray, v...)
	default:
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.InvalidTypeForFilesToAddError],
		})

		os.Exit(1)
	}

	finalContent := strings.Join(trimmedIgnoredFilesArray, "\n")

	writeFile(constants.PATH_DIR_NPMIGNORE, []byte(finalContent))

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "\".npmignore\" file modified successfully",
	})
}
