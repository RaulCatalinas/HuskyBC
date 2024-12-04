package handlers

import (
	"os"
	"time"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	errorMessages "github.com/RaulCatalinas/HuskyBC/internal/error_messages"
	userInput "github.com/RaulCatalinas/HuskyBC/internal/user_input"
	"github.com/RaulCatalinas/HuskyBC/internal/utils"
	"github.com/toqueteos/webbrowser"
)

func HandlerOptionCollaborate() {
	utils.WriteMessage(utils.WriteMessageProps{
		Type:    enums.MessageTypeInfo,
		Message: "Opening the GitHub repository...",
	})

	time.Sleep(5 * time.Millisecond)

	err := webbrowser.Open(constants.REPOSITORY)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.GitHubRepoOpenError],
		})

		os.Exit(1)
	}
}

func HandlerOptionBuild() {
	utils.CheckinFolderOrFile(constants.PATH_PACKAGE_JSON, false)

	packageManagerToUse := userInput.GetPackageManager()
	useCommitlint := userInput.AddCommitlint()
	shouldPublishToNpm := userInput.ShouldPublishToNpm()

	utils.GenerateHuskyConfig(
		utils.Props{
			PackageManagerToUse: packageManagerToUse,
			PackageJsonPath:     constants.PATH_PACKAGE_JSON,
			UseCommitlint:       useCommitlint,
			ShouldPublishToNpm:  shouldPublishToNpm,
		},
	)

	if useCommitlint {
		utils.GenerateCommitlintConfig(
			utils.CommitlintProps{
				PackageManagerToUse: packageManagerToUse,
				PackageJsonPath:     constants.PATH_PACKAGE_JSON,
				ShouldPublishToNpm:  shouldPublishToNpm,
			},
		)
	}

	utils.WriteMessage(utils.WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: "All tasks were completed",
	})
}
