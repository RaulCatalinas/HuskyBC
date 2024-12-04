package error_messages

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

var PROCESS_ERROR_MESSAGES = map[enums.ProcessError]string{
	enums.JsonUnmarshalError: "Replace this with an error message that's descriptive and easily understandable by any developer.",

	enums.JsonMarshalError: "Replace this with an error message that's descriptive and easily understandable by any developer.",

	enums.GetWorkingDirectoryError: "Replace this with an error message that's descriptive and easily understandable by any developer.",

	enums.InvalidTypeForFilesToAddError: "Replace this with an error message that's descriptive and easily understandable by any developer.",

	enums.CheckingFolderOrFileError: "Replace this with an error message that's descriptive and easily understandable by any developer.",

	enums.DependenciesError: generateErrorMessage("installing dependencies"),

	enums.GitHubRepoOpenError: generateErrorMessage("opening the GitHub repository in a new browser tab"),
}
