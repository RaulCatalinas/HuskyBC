package handlers

import (
	"os"
	"sync"

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

	if err := webbrowser.Open(constants.REPOSITORY); err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.PROCESS_ERROR_MESSAGES[enums.GitHubRepoOpenError],
		})

		os.Exit(1)
	}
}

func HandlerOptionBuild() {
	packageManagerToUse := userInput.GetPackageManager()
	useCommitlint := userInput.AddCommitlint()
	shouldPublishToNpm := userInput.ShouldPublishToNpm()

	utils.CreateEmptyJsonFileIfNotExists(constants.PATH_PACKAGE_JSON)

	dependenciesToInstall := []string{"husky"}

	if useCommitlint {
		dependenciesToInstall = append(
			dependenciesToInstall,
			"lint-staged",
			"@commitlint/cli",
			"@commitlint/config-conventional",
		)
	}

	utils.InstallDependencies(utils.InstallProps{
		PackageManagerToUse: packageManagerToUse,
		PackagesToInstall:   dependenciesToInstall,
	})

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		utils.GenerateHuskyConfig(
			utils.Props{
				PackageManagerToUse: packageManagerToUse,
				PackageJsonPath:     constants.PATH_PACKAGE_JSON,
				UseCommitlint:       useCommitlint,
				ShouldPublishToNpm:  shouldPublishToNpm,
			},
		)
	}()

	if useCommitlint {
		wg.Add(1)

		go func() {
			defer wg.Done()

			utils.GenerateCommitlintConfig(
				utils.CommitlintProps{
					PackageManagerToUse: packageManagerToUse,
					PackageJsonPath:     constants.PATH_PACKAGE_JSON,
					ShouldPublishToNpm:  shouldPublishToNpm,
				},
			)
		}()
	}

	wg.Wait()

	utils.WriteMessage(utils.WriteMessageProps{
		Type:    enums.MessageTypeSuccess,
		Message: "All tasks were completed",
	})
}
