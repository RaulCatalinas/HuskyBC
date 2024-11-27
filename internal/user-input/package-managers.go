package userinput

import (
	"fmt"
	"os"

	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	"github.com/RaulCatalinas/HuskyBC/internal/utils"

	"github.com/AlecAivazis/survey/v2"
)

func GetPackageManager() enums.PackageManager {
	var questions = []*survey.Question{
		{
			Name: "packageManger",
			Prompt: &survey.Select{
				Options: utils.GetPackageManagersAsStrings(),
				Message: "Which package manager do you wanna use?",
				Default: "npm", // <- If we use the enum here the library gives an error.
			},
		},
	}

	answers := struct {
		PackageManager string `survey:"packageManger"`
	}{}

	utils.WriteMessage(utils.WriteMessageProps{
		Type:    enums.MessageTypeWarning,
		Message: "When using Bun as a package manager it's possible that the dependencies aren't installed correctly, we're working on it.",
	})

	fmt.Println()

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: utils.GetErrorMessage("PackageManagerSelection"),
		})

		os.Exit(1)
	}

	return enums.PackageManager(answers.PackageManager)
}
