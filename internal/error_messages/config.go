package error_messages

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

var CONFIG_ERROR_MESSAGES = map[enums.ConfigError]string{
	enums.HuskyError:             generateErrorMessage("configuring Husky"),
	enums.CommitLintConfigError:  generateErrorMessage("configuring commitlint"),
	enums.ConfigFilesCreateError: generateErrorMessage("creating configuration files"),
}
