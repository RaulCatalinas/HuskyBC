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
	WriteMessage(WriteMessageProps{
		Type:    "config",
		Message: "Generating Husky's Configuration...",
	})

	InstallDependencies(InstallProps{
		PackageManagerToUse: props.PackageManagerToUse,
		PackagesToInstall:   []string{"husky"},
	})

	createHuskyConfigFiles(props.PackageManagerToUse, props.UseCommitlint)

	addNecessaryScriptsToPakageJson(
		props.PackageJsonPath,
		props.PackageManagerToUse,
		props.ShouldPublishToNpm,
	)

	if props.ShouldPublishToNpm {
		modifyNpmIgnore(".husky")
	}

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "Husky's configuration generated successfully",
	})
}

func createHuskyConfigFiles(packageManagerToUse types.PackageManager, useCommitlint bool) {
	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: "Creating configuration file...",
	})

	CheckinFolderOrFile(constants.PATH_DIR_HUSKY, true)

	var preCommitFileValue string

	if useCommitlint {
		preCommitFileValue = constants.LINT_STAGED_CONFIG[types.PackageManager(packageManagerToUse)]
	} else {
		preCommitFileValue = string(packageManagerToUse) + "test"
	}

	writeFile(constants.PATH_DIR_HUSKY+"/pre-commit", []byte(preCommitFileValue))

	WriteMessage(WriteMessageProps{
		Type:    "success",
		Message: "Configuration file (pre-commit) created successfully",
	})
}

type packageJsonScript struct {
	Key   string
	Value string
}

func addNecessaryScriptsToPakageJson(
	packageJsonPath string,
	packageManagerToUse types.PackageManager,
	shouldPublishToNpm bool,
) {
	WriteMessage(WriteMessageProps{
		Type:    "info",
		Message: "Adding necessary scripts to package.json...",
	})

	var huskyScriptsForYarn interface{}
	var scriptsToAdd interface{}

	if shouldPublishToNpm {
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

	addScript(addScriptProps{
		ScriptsToAdd:    scriptsToAdd,
		PackageJsonPath: packageJsonPath,
	})
}
