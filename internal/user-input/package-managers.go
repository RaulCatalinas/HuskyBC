package userinput

import (
	"fmt"
	"os"

	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/types"
	"github.com/RaulCatalinas/HuskyBC/internal/utils"

	"github.com/AlecAivazis/survey/v2"
)

func GetPackageManager() types.PackageManager {
	var questions = []*survey.Question{
		{
			Name: "packageManger",
			Prompt: &survey.Select{
				Options: constants.PACKAGE_MANAGERS,
				Message: "Which package manager do you wanna use?",
				Default: "npm",
			},
		},
	}

	answers := struct {
		PackageManager string `survey:"packageManger"`
	}{}

	utils.WriteMessage(utils.WriteMessageProps{
		Type:    "warning",
		Message: "When using Bun as a package manager it's possible that the dependencies aren't installed correctly, we're working on it.",
	})

	fmt.Println()

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    "error",
			Message: utils.GetErrorMessage("PackageManagerSelection"),
		})

		os.Exit(1)
	}

	return types.PackageManager(answers.PackageManager)
}
