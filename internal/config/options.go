package config

import (
	"os"

	"github.com/RaulCatalinas/HuskyBC/internal/cli"
	"github.com/RaulCatalinas/HuskyBC/internal/options"
)

func ConfigureOptions() {
	options := options.GetOptions()

	if len(os.Args) != 2 {
		cli.ShowHelp(options)

		os.Exit(0)
	}

	for _, option := range options {
		if os.Args[1] == option.Name || os.Args[1] == option.Alias {
			option.Handler()

			return
		}

		cli.ShowHelp(options)
	}
}
