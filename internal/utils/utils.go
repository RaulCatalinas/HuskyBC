package utils

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	errorMessages "github.com/RaulCatalinas/HuskyBC/internal/error_messages"
)

func CheckIfExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func createEmptyFile(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.FILE_ERROR_MESSAGES[enums.CreateEmptyFileError],
		})

		os.Exit(1)
	}
	defer file.Close()
}

func createFolder(name string) {
	os.Mkdir(name, 0750)

	message := "Created folder " + name

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: message,
	})
}

func createFolderOrFile(path string, isFolder bool) {
	if isFolder {
		createFolder(path)
	} else {
		createEmptyFile(path)
	}
}

func CreateEmptyJsonFileIfNotExists(fileName string) {
	exist := CheckIfExists(fileName)

	if !exist {
		createEmptyFile(fileName)

		writeFile(fileName, []byte("{}"))
	}
}

func readFile(filename string) []byte {
	dataByte, err := os.ReadFile(filename)

	if err != nil {
		errorMessage := errorMessages.FILE_ERROR_MESSAGES[enums.ReadFileError]

		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: strings.Replace(errorMessage, "{fileName}", filename, -1),
		})

		os.Exit(1)
	}

	return dataByte
}

func writeFile(filename string, newData []byte) {
	err := os.WriteFile(filename, newData, 0644)

	if err != nil {
		errorMessage := errorMessages.FILE_ERROR_MESSAGES[enums.WriteFileError]

		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: strings.Replace(errorMessage, "{fileName}", filename, -1),
		})
	}
}

func jsonMarshalIndent(packageJsonObj map[string]interface{}) []byte {
	newData, err := json.MarshalIndent(packageJsonObj, "", "  ")

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.JsonMarshalError],
		})

		os.Exit(1)
	}

	return newData
}

func jsonUnmarshal(data []byte) map[string]interface{} {
	var packageJsonObj map[string]interface{}

	err := json.Unmarshal(data, &packageJsonObj)

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.JsonUnmarshalError],
		})

		os.Exit(1)
	}

	return packageJsonObj
}

func readAndUnmarshalPackageJson(packageJsonPath string) map[string]interface{} {
	dataByte := readFile(packageJsonPath)

	packageJsonObj := jsonUnmarshal(dataByte)

	return packageJsonObj
}

func jsonMarshalIndentAndWriteFile(packageJsonObj map[string]interface{}, filename string) {
	newData := jsonMarshalIndent(packageJsonObj)

	writeFile(filename, newData)
}
