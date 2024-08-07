package utils

import (
	"fmt"
	"os"
	"strings"
)

func modifyNpmIgnore(filesToAdd interface{}) {
	defer func() {
		if r := recover(); r != nil {
			WriteMessage(WriteMessageProps{
				Type:    "error",
				Message: GetErrorMessage("NpmIgnoreWrite"),
			})

			os.Exit(1)
		}
	}()

	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: "Writing in the file \".npmignore\"...",
	})

	dir, err := os.Getwd()

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    "error",
			Message: GetErrorMessage("GetWorkingDirectory"),
		})

		os.Exit(1)
	}

	npmIgnoreFilePath := fmt.Sprintf("%s/.npmignore", dir)

	if !exists(npmIgnoreFilePath) {
		createEmptyFile(".npmignore")
	}

	data := readFile(npmIgnoreFilePath)

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
				Type:    "error",
				Message: GetErrorMessage("InvalidTypeForFilesToAdd"),
			})

			os.Exit(1)
		}

		writeFile(npmIgnoreFilePath, []byte(filesToWrite))

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
			Type:    "error",
			Message: GetErrorMessage("InvalidTypeForFilesToAdd"),
		})

		os.Exit(1)
	}

	finalContent := strings.Join(trimmedIgnoredFilesArray, "\n")

	writeFile(npmIgnoreFilePath, []byte(finalContent))

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "\".npmignore\" file modified successfully",
	})
}
