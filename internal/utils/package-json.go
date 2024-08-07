package utils

import (
	"os"
)

type addScriptProps struct {
	ScriptsToAdd    interface{} // Puede ser packageJsonScript o []packageJsonScript
	PackageJsonPath string
}

func existsSection(packageJsonPath string, sectionToCheck string) bool {
	packageJsonObj := readAndUnmarshalPackageJson(packageJsonPath)

	_, exists := packageJsonObj[sectionToCheck]
	return exists
}

func createEmptySection(packageJsonPath string, sectionToCreate string) {
	packageJsonObj := readAndUnmarshalPackageJson(packageJsonPath)

	packageJsonObj[sectionToCreate] = map[string]interface{}{}

	jsonMarshalIndentAndWriteFile(packageJsonObj, packageJsonPath)
}

func addScript(props addScriptProps) {
	packageJsonPath := props.PackageJsonPath
	scriptsToAdd := props.ScriptsToAdd

	defer func() {
		if r := recover(); r != nil {
			WriteMessage(WriteMessageProps{
				Type:    "error",
				Message: GetErrorMessage("AddScript"),
			})

			os.Exit(1)
		}
	}()

	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: "Modifying package.json...",
	})

	if !existsSection(packageJsonPath, "scripts") {
		createEmptySection(packageJsonPath, "scripts")
	}

	packageJsonObj := readAndUnmarshalPackageJson(packageJsonPath)

	scriptsSection := packageJsonObj["scripts"].(map[string]interface{})

	switch v := scriptsToAdd.(type) {
	case []packageJsonScript:
		for _, script := range v {
			scriptsSection[script.Key] = script.Value
		}
	case packageJsonScript:
		scriptsSection[v.Key] = v.Value
	default:
		WriteMessage(WriteMessageProps{
			Type:    "error",
			Message: GetErrorMessage("InvalidTypeForScriptsToAdd"),
		})
		os.Exit(1)
	}

	jsonMarshalIndentAndWriteFile(packageJsonObj, packageJsonPath)

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "Package.json modified successfully",
	})
}
