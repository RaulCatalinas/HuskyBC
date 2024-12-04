package user_input

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	errorMessages "github.com/RaulCatalinas/HuskyBC/internal/error_messages"
	"github.com/RaulCatalinas/HuskyBC/internal/utils"
)

func AddCommitlint() bool {
	var questions = []*survey.Question{
		{
			Name: "addCommitlint",
			Prompt: &survey.Confirm{
				Message: "Do you wanna add commitlint?",
				Default: true,
			},
		},
	}

	answers := struct {
		AddCommitlint bool `survey:"addCommitlint"`
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		utils.WriteMessage(utils.WriteMessageProps{
			Type:    enums.MessageTypeError,
			Message: errorMessages.USER_INPUT_ERROR_MESSAGES[enums.AddCommitLintError],
		})

		os.Exit(1)
	}

	return answers.AddCommitlint
}
