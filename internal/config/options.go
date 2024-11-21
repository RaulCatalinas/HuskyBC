package config

import (
	"os"

	"github.com/RaulCatalinas/HuskyBC/internal/cli"
	"github.com/RaulCatalinas/HuskyBC/internal/constants"
)

func ConfigureOptions() {
	if len(os.Args) != 2 {
		cli.ShowHelp()

		os.Exit(0)
	}

	if len(os.Args) != 2 {
		cli.ShowHelp()
		os.Exit(0)
	}

	for _, option := range constants.Options {
		if os.Args[1] == option.Name || os.Args[1] == option.Alias {
			option.Handler()

			return
		}

		cli.ShowHelp()
	}
}
