package constants

import "github.com/RaulCatalinas/HuskyBC/internal/types"

var Options = []types.Option{
	{
		Name:        "--version",
		Alias:       "-v",
		Description: "Output the version number",
	},
	{
		Name:        "--collaborate",
		Alias:       "-co",
		Description: "Open GitHub repository for collaboration",
	},
	{
		Name:        "--build",
		Alias:       "-b",
		Description: "Start Husky's configuration",
	},
	{
		Name:        "--help",
		Alias:       "-h",
		Description: "Display help for command",
	},
}
