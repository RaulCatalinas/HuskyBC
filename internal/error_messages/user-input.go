package error_messages

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

var USER_INPUT_ERROR_MESSAGES = map[enums.UserInputError]string{
	enums.AddCommitLintError: generateErrorMessage("asking if you want to add commitlint"),

	enums.PackageManagerSelectionError: generateErrorMessage("selecting the package manager"),

	enums.ShouldPublishToNpmError: generateErrorMessage("confirming npm publication"),
}
