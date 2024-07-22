package userinput

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
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
			Type:    "error",
			Message: utils.GetErrorMessage("AddCommitLint"),
		})

		os.Exit(1)
	}

	return answers.AddCommitlint
}
