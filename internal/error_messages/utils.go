package error_messages

import "github.com/RaulCatalinas/HuskyBC/internal/constants"

func generateErrorMessage(action string) string {
	return "An error occurred while " + action + ". " + constants.BASE_ERROR_MESSAGE
}
