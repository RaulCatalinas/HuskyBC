package userinput

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	"github.com/RaulCatalinas/HuskyBC/internal/utils"
)

func ShouldPublishToNpm() bool {
	var questions = []*survey.Question{
		{
			Name: "shouldPublishToNpm",
			Prompt: &survey.Confirm{
				Message: "Will you publish it on npm?",
				Default: false,
			},
		},
	}

	answers := struct {
		ShouldPublishToNpm bool `survey:"shouldPublishToNpm"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: utils.GetErrorMessage("ShouldPublishToNpm"),
		})

		os.Exit(1)
	}

	return answers.ShouldPublishToNpm
}
