package utils

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/RaulCatalinas/HuskyBC/internal/enums"
)

func exists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !os.IsNotExist(err)
}

func createEmptyFile(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		WriteMessage(WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: GetErrorMessage("CreateEmptyFile"),
		})

		os.Exit(1)
	}
	defer file.Close()
}

func readFile(filename string) []byte {
	dataByte, err := os.ReadFile(filename)

	if err != nil {
		errorMessage := GetErrorMessage("ReadFile")

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
		errorMessage := GetErrorMessage("WriteFile")

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
			Message: GetErrorMessage("JsonMarshal"),
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
			Message: GetErrorMessage("JsonUnmarshal"),
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
