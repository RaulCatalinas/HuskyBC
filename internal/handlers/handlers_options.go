package handlers

import (
	"os"
	"time"

	"github.com/toqueteos/webbrowser"
	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	userinput "github.com/RaulCatalinas/HuskyBC/internal/user-input"
	"github.com/RaulCatalinas/HuskyBC/internal/utils"
)

func HandlerOptionCollaborate() {
	utils.WriteMessage("info", "Opening the GitHub repository...")
	time.Sleep(5 * time.Millisecond)
	err := webbrowser.Open(constants.REPOSITORY)
	if err != nil {
		utils.WriteMessage("error", utils.GetErrorMessage("Open webbrowser"))
		os.Exit(1)
	}
}

func HandlerOptionBuild() {
	utils.CheckinFolderOrFile(constants.PATH_PACKAGE_JSON, false)

	packageManagerToUse := userinput.GetPackageManager()
	shouldPublishToNpm := userinput.ShouldPublishToNPM()
	useCommitlint := userinput.AddCommitlint()

	utils.GenerateHuskyConfig(
		utils.Props{
			PackageManagerToUse: packageManagerToUse,
			PackageJsonPath:     constants.PATH_PACKAGE_JSON,
			UseCommitlint:       useCommitlint,
			ShouldPublishToNpm:  shouldPublishToNpm,
		})
	if useCommitlint {
		utils.GenerateCommitlintConfig(
			utils.CommitlintProps{
				PackageManagerToUse: packageManagerToUse,
				PackageJsonPath:     constants.PATH_PACKAGE_JSON,
				ShouldPublishToNpm:  shouldPublishToNpm,
			})
	}
	utils.WriteMessage("success", "All tasks were completed")
}
