package utils

import (
	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
)

type Props struct {
	PackageManagerToUse enums.PackageManager
	PackageJsonPath     string
	UseCommitlint       bool
	ShouldPublishToNpm  bool
}

func GenerateHuskyConfig(props Props) {
	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeConfig,
		Message: "Generating Husky's Configuration...",
	})

	InstallDependencies(InstallProps{
		PackageManagerToUse: props.PackageManagerToUse,
		PackagesToInstall:   []string{"husky"},
	})

	createHuskyConfigFiles(props.PackageManagerToUse, props.UseCommitlint)

	addNecessaryScriptsToPackageJson(
		props.PackageJsonPath,
		props.PackageManagerToUse,
		props.ShouldPublishToNpm,
	)

	if props.ShouldPublishToNpm {
		modifyNpmIgnore(".husky")
	}

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: "Husky's configuration generated successfully",
	})
}

func createHuskyConfigFiles(packageManagerToUse enums.PackageManager, useCommitlint bool) {
	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: "Creating configuration file...",
	})

	CheckinFolderOrFile(constants.PATH_DIR_HUSKY, true)

	var preCommitFileValue string

	if useCommitlint {
		preCommitFileValue = constants.LINT_STAGED_CONFIG[enums.PackageManager(packageManagerToUse)]
	} else {
		preCommitFileValue = string(packageManagerToUse) + "test"
	}

	writeFile(constants.PATH_DIR_HUSKY+"/pre-commit", []byte(preCommitFileValue))

	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: "Configuration file (pre-commit) created successfully",
	})
}

type packageJsonScript struct {
	Key   string
	Value string
}

func addNecessaryScriptsToPackageJson(
	packageJsonPath string,
	packageManagerToUse enums.PackageManager,
	shouldPublishToNpm bool,
) {
	WriteMessage(WriteMessageProps{
		Type:    enums.MessageTypeInfo,
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
