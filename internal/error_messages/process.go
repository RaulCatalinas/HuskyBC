package error_messages

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

var PROCESS_ERROR_MESSAGES = map[enums.ProcessError]string{
	enums.JsonUnmarshalError: generateErrorMessage("unmarshalling JSON data"),

	enums.JsonMarshalError: generateErrorMessage("marshalling JSON data with indentation"),

	enums.InvalidTypeForFilesToAddError: generateErrorMessage("adding files to .npmignore with an invalid type"),

	enums.DependenciesError: generateErrorMessage("installing dependencies"),

	enums.GitHubRepoOpenError: generateErrorMessage("opening the GitHub repository in a new browser tab"),
}
