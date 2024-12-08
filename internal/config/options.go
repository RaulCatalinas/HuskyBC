package config

import (
	"fmt"
	"os"

	"github.com/RaulCatalinas/HuskyBC/internal/cli"
	"github.com/RaulCatalinas/HuskyBC/internal/enums"
	"github.com/RaulCatalinas/HuskyBC/internal/options"
	"github.com/RaulCatalinas/HuskyBC/internal/types"
	"github.com/RaulCatalinas/HuskyBC/internal/utils"
)

func buildOptionMap(optionsList []types.Option) map[string]func() {
	optionMap := make(map[string]func())

	for _, option := range optionsList {
		optionMap[option.Name] = option.Handler
		optionMap[option.Alias] = option.Handler
	}

	return optionMap
}

func getOptionHandler(optionName string, availableOptions []types.Option) func() {
	optionMap := buildOptionMap(availableOptions)

	if handler, exists := optionMap[optionName]; exists {
		return handler
	}

	return nil
}

func handleInvalidOption(availableOptions []types.Option) {
	utils.WriteMessage(utils.WriteMessageProps{
		Type:    enums.MessageTypeError,
		Message: "The option you've tried to execute doesn't exist",
	})

	fmt.Println()

	showHelpAndExit(availableOptions, 1)
}

func showHelpAndExit(options []types.Option, exitCode int) {
	cli.ShowHelp(options)

	os.Exit(exitCode)
}

func ConfigureOptions() {
	availableOptions := options.GetOptions()

	if len(os.Args) != 2 {
		showHelpAndExit(availableOptions, 0)
	}

	optionName := os.Args[1]

	if handler := getOptionHandler(optionName, availableOptions); handler != nil {
		handler()

		return
	}

	handleInvalidOption(availableOptions)
}
