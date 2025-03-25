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

type optionManager struct {
	availableOptions []types.Option
	optionHandlers   map[string]func()
}

func NewOptionManager() *optionManager {
	availableOptions := options.GetOptions()

	return &optionManager{
		availableOptions: availableOptions,
		optionHandlers:   buildOptionHandlerMap(availableOptions),
	}
}

func buildOptionHandlerMap(optionsList []types.Option) map[string]func() {
	optionHandlers := make(map[string]func(), len(optionsList)*2)

	for _, option := range optionsList {
		optionHandlers[option.Name] = option.Handler
		optionHandlers[option.Alias] = option.Handler
	}

	return optionHandlers
}

func (optionManager *optionManager) getOptionHandler(optionName string) func() {
	return optionManager.optionHandlers[optionName]
}

func (optionManager *optionManager) handleInvalidOption() {
	utils.WriteMessage(utils.WriteMessageProps{
		Type:    enums.MessageTypeError,
		Message: "The option you've tried to execute doesn't exist",
	})

	fmt.Println()
	optionManager.showHelpAndExit(1)
}

func (optionManager *optionManager) showHelpAndExit(exitCode int) {
	cli.ShowHelp(optionManager.availableOptions)
	os.Exit(exitCode)
}

func (optionManager *optionManager) ExecuteSelectedOption() {
	if len(os.Args) != 2 {
		optionManager.showHelpAndExit(0)
	}

	optionName := os.Args[1]

	handler := optionManager.getOptionHandler(optionName)

	if handler != nil {
		handler()
		return
	}

	optionManager.handleInvalidOption()
}
