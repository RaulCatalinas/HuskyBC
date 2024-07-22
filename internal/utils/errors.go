package utils

import "github.com/RaulCatalinas/HuskyBC/internal/constants"

func GetErrorMessage(error string) string {
	return constants.ERROR_MESSAGES[error]
}
