package constants

var ERROR_MESSAGES = map[string]string{
	"NotFound": "The {fileName} file wasn't found in the current directory.",

	"Default": "Something went wrong, please try again later, if the error persists please report it on " + ISSUES + ".",

	"Dependencies": "An error occurred while installing dependencies, please try again later, if the error persists please report it on " + ISSUES + ".",

	"Husky": "An error has occurred during the Husky configuration process, please try again later, if the error persists please report it on " + ISSUES + ".",

	"CommitLintConfig": "An error has occurred during the commitlint configuration process, please try again later, if the error persists please report it on " + ISSUES + ".",

	"AddCommitLint": "An error occurred while asking if you wanna add commitlint, please try again later, if the error persists please report it on " + ISSUES + ".",

	"PackageManagerSelection": "An error occurred while selecting the package manager, please try again later, if the error persists, please report it on " + ISSUES + ".",

	"CreateEmptyFile": "An error occurred while creating the empty file. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"CreateFolder": "An error occurred while creating the folder: {folderName}, please try again later, if the error persists, please report it on " + ISSUES + ".",

	"ReadFile": "An error has occurred while reading the {fileName} file, please try again later, if the error persists, please report it on " + ISSUES + ".",

	"WriteFile": "An error has occurred while writing to the {fileName} file, please try again later, if the error persists, please report it on " + ISSUES + ".",

	"ConfigFilesCreate": "An error occurred while creating configuration files. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"ShouldPublishToNpm": "An error occurred while confirming npm publication. Please try again later, if the error persists, please report it on" + ISSUES + ".",

	"NpmIgnoreWrite": "An error occurred while writing to the '.npmignore' file. Please try again later, if the error persists please report it on " + ISSUES + ".",

	"GitHubRepoOpen": "An error occurred while opening the GitHub repository in a new browser tab. Please try again later, if the error persists please report it on" + ISSUES + ".",

	"JsonUnmarshal": "Replace this with an error message that's descriptive and easily understandable by any developer.",

	"JsonMarshal": "Replace this with an error message that's descriptive and easily understandable by any developer.",

	"GetWorkingDirectory": "Replace this with an error message that's descriptive and easily understandable by any developer.",

	"InvalidTypeForFilesToAdd": "Replace this with an error message that's descriptive and easily understandable by any developer.",

	"CheckingFolderOrFile": "Replace this with an error message that's descriptive and easily understandable by any developer.",
}
