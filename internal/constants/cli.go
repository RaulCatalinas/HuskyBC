package constants

import "github.com/RaulCatalinas/HuskyBC/internal/types"

var Options = []types.Option{
	{
		Option:      "--version",
		Alias:       "-v",
		Description: "Output the version number",
	},
	{
		Option:      "--collaborate",
		Alias:       "-co",
		Description: "Open GitHub repository for collaboration",
	},
	{
		Option:      "--build",
		Alias:       "-b",
		Description: "Start Husky's configuration",
	},
	{
		Option:      "--help",
		Alias:       "-h",
		Description: "Display help for command",
	},
}
