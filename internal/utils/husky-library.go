package utils

import (
	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/types"
)

type Props struct {
	PackageManagerToUse types.PackageManager
	PackageJsonPath     string
	UseCommitlint       bool
	ShouldPublishToNpm  bool
}

func GenerateHuskyConfig(props Props) {
	WriteMessage("info", "Generating Husky's Configuration...")

	InstallDependencies(InstallProps{PackageManagerToUse: props.PackageManagerToUse, PackagesToInstall: []string{"husky"}})

	createHuskyConfigFiles(props.PackageManagerToUse, props.UseCommitlint)

	addNecessaryScriptsToPakageJson(props.PackageJsonPath, props.PackageManagerToUse, props.ShouldPublishToNpm)

	if props.ShouldPublishToNpm {
		modifyNpmIgnore(".husky")
	}

	WriteMessage("success", "Husky's configuration generated successfully")
}

func createHuskyConfigFiles(packageManagerToUse types.PackageManager, useCommitlint bool) {
	WriteMessage("info", "Creating configuration file...")

	CheckinFolderOrFile(constants.PATH_DIR_HUSKY, true)

	var preCommitFileValue string
	if useCommitlint {
		preCommitFileValue = constants.LINT_STAGED_CONFIG[types.PackageManager(packageManagerToUse)]
	} else {
		preCommitFileValue = string(packageManagerToUse) + "test"
	}

	writeFile(constants.PATH_DIR_HUSKY+"/pre-commit", []byte(preCommitFileValue))

	WriteMessage("info", "Configuration file (pre-commit) created successfully")
}

type packageJsonScript struct {
	Key   string
	Value string
}

func addNecessaryScriptsToPakageJson(packaJsonPath string, packageManagerToUse types.PackageManager, sholdPublishToNpm bool) {
	WriteMessage("info", "Add necessary scripts to "+packaJsonPath)

	var huskyScriptsForYarn interface{}
	var scriptsToAdd interface{}

	if sholdPublishToNpm {
		huskyScriptsForYarn = packageJsonScript{Key: "postintall", Value: "husky"}
	} else {
		huskyScriptsForYarn = []packageJsonScript{
			{Key: "postinstall", Value: "husky"},
			{Key: "prepack", Value: "pinst --disable"},
			{Key: "postpack", Value: "pinst --enable"},
		}
	}

	if packageManagerToUse != "yarn" {
		scriptsToAdd = packageJsonScript{Key: "prepare", Value: "husky"}
	} else {
		scriptsToAdd = huskyScriptsForYarn
	}

	addScript(addScriptProps{ScriptsToAdd: scriptsToAdd, PackageJsonPath: packaJsonPath})
}
