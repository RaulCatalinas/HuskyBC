package config

import (
	"os"

	"github.com/RaulCatalinas/HuskyBC/internal/cli"
	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/handlers"
)

func ConfigureOptions() {
	if len(os.Args) != 2 {
		cli.ShowHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case constants.Options[1].Option, constants.Options[1].Alias:
		handlers.HandlerOptionCollaborate()
	case constants.Options[2].Option, constants.Options[2].Alias:
		handlers.HandlerOptionBuild()
	case constants.Options[0].Option, constants.Options[0].Alias:
		cli.ShowVersion()
	case constants.Options[3].Option, constants.Options[3].Alias:
		cli.ShowHelp()
	default:
		cli.ShowHelp()
	}
}
