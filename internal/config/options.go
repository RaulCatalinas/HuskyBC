package config

import (
	"os"

	"github.com/RaulCatalinas/HuskyBC/internal/cli"
	"github.com/RaulCatalinas/HuskyBC/internal/constants"
	"github.com/RaulCatalinas/HuskyBC/internal/handlers"
	"github.com/RaulCatalinas/HuskyBC/internal/types"
)

func ConfigureOptions() {
	if len(os.Args) != 2 {
		cli.ShowHelp()

		os.Exit(0)
	}

	// TODO: Move this constant to a more appropriate place
	var commands = []types.Command{
		{
			Name:    constants.Options[0].Name,
			Alias:   constants.Options[0].Alias,
			Handler: cli.ShowVersion,
		},
		{
			Name:    constants.Options[1].Name,
			Alias:   constants.Options[1].Alias,
			Handler: handlers.HandlerOptionCollaborate,
		},
		{
			Name:    constants.Options[2].Name,
			Alias:   constants.Options[2].Alias,
			Handler: handlers.HandlerOptionBuild,
		},
		{
			Name:    constants.Options[3].Name,
			Alias:   constants.Options[3].Alias,
			Handler: cli.ShowHelp,
		},
	}

	if len(os.Args) != 2 {
		cli.ShowHelp()
		os.Exit(0)
	}

	for _, cmd := range commands {
		if os.Args[1] == cmd.Name || os.Args[1] == cmd.Alias {
			cmd.Handler()

			return
		}

		cli.ShowHelp()
	}
}
