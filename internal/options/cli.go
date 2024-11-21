package options

import (
	"github.com/RaulCatalinas/HuskyBC/internal/cli"
	"github.com/RaulCatalinas/HuskyBC/internal/handlers"
	"github.com/RaulCatalinas/HuskyBC/internal/types"
)

var Options = []types.Option{
	{
		Name:        "--version",
		Alias:       "-v",
		Description: "Output the version number",
		Handler:     cli.ShowVersion,
	},
	{
		Name:        "--collaborate",
		Alias:       "-co",
		Description: "Open GitHub repository for collaboration",
		Handler:     handlers.HandlerOptionCollaborate,
	},
	{
		Name:        "--build",
		Alias:       "-b",
		Description: "Start Husky's configuration",
		Handler:     handlers.HandlerOptionBuild,
	},
	{
		Name:        "--help",
		Alias:       "-h",
		Description: "Display help for command",
		Handler:     cli.ShowHelp,
	},
}
