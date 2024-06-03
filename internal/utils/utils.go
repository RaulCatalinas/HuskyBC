package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetErrorMessage(context string) string {
	return fmt.Sprintf("Error in %s", context)
}

func WriteMessage(messageType string, message string) {
	log.Printf("[%s] %s", messageType, message)
}

func exists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !os.IsNotExist(err)
}

func createEmptyFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		WriteMessage("error", GetErrorMessage("CreateFile"))
		os.Exit(1)
	}
	defer file.Close()
}

func readFile(filename string) []byte {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		WriteMessage("error", GetErrorMessage("Read file "+filename))
		os.Exit(1)
	}

	return dataByte
}

func writeFile(filename string, newData []byte) {
	err := os.WriteFile(filename, newData, 0644)
	if err != nil {
		WriteMessage("error", GetErrorMessage("Write file "+filename))
	}
}

func jsonMarshalIndent(packageJsonObj map[string]interface{}) []byte {
	newData, err := json.MarshalIndent(packageJsonObj, "", "  ")
	if err != nil {
		WriteMessage("error", GetErrorMessage("Json marshal indent"))
		os.Exit(1)
	}

	return newData
}

func jsonUnmarshal(data []byte) map[string]interface{} {
	var packageJsonObj map[string]interface{}

	err := json.Unmarshal(data, &packageJsonObj)
	if err != nil {
		WriteMessage("error", GetErrorMessage("Json unmarshal"))
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
