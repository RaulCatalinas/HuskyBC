package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	errorMessages "github.com/RaulCatalinas/HuskyBC/internal/error_messages"
)

func handlePanic() {
	if r := recover(); r != nil {
		logError(errorMessages.FILE_ERROR_MESSAGES[enums.NpmIgnoreWriteError])
		os.Exit(1)
	}
}

func logInfo(message string) {
	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: message,
	})
}

func logSuccess(message string) {
	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: message,
	})
}

func logError(message string) {
	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeError,
		Message: message,
	})
}

func parseFilesToAdd(filesToAdd interface{}) []string {
	switch v := filesToAdd.(type) {
	case string:
		return []string{strings.TrimSpace(v)}
	case []string:
		cleaned := make([]string, 0, len(v))

		for _, file := range v {
			trimmed := strings.TrimSpace(file)

			if trimmed != "" {
				cleaned = append(cleaned, trimmed)
			}
		}

		return cleaned
	default:
		logError(errorMessages.PROCESS_ERROR_MESSAGES[enums.InvalidTypeForFilesToAddError])

		os.Exit(1)
	}

	return nil
}

func mergeFiles(existingFiles, newFiles []string) []string {
	fileSet := make(map[string]struct{}, len(existingFiles)+len(newFiles))

	for _, file := range existingFiles {
		fileSet[file] = struct{}{}
	}

	for _, file := range newFiles {
		fileSet[file] = struct{}{}
	}

	merged := make([]string, 0, len(fileSet))

	for file := range fileSet {
		merged = append(merged, file)
	}

	return merged
}

func readNpmIgnore(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}

		return nil, err
	}

	defer file.Close()

	var files []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line != "" {
			files = append(files, line)
		}
	}

	return files, scanner.Err()
}

func writeNpmIgnore(filePath string, content []string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range content {
		_, err := writer.WriteString(line + "\n")

		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func modifyNpmIgnore(filesToAdd interface{}) {
	defer handlePanic()

	logInfo("Writing to the \".npmignore\" file...")

	CreateFolderOrFileIfNotExists(constants.PATH_DIR_NPMIGNORE, false)

	existingFiles, err := readNpmIgnore(constants.PATH_DIR_NPMIGNORE)

	if err != nil {
		logError(errorMessages.FILE_ERROR_MESSAGES[enums.NpmIgnoreWriteError])

		os.Exit(1)
	}

	filesToWrite := parseFilesToAdd(filesToAdd)
	finalContent := mergeFiles(existingFiles, filesToWrite)

	err = writeNpmIgnore(constants.PATH_DIR_NPMIGNORE, finalContent)

	if err != nil {
		logError(errorMessages.FILE_ERROR_MESSAGES[enums.NpmIgnoreWriteError])
		os.Exit(1)
	}

	logSuccess("\".npmignore\" file modified successfully")
}
