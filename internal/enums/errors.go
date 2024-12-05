package enums

type FileError string
type ConfigError string
type ProcessError string
type UserInputError string

const (
	NotFoundError        FileError = "NotFound"
	ReadFileError        FileError = "ReadFile"
	WriteFileError       FileError = "WriteFile"
	CreateEmptyFileError FileError = "CreateEmptyFile"
	CreateFolderError    FileError = "CreateFolder"
	NpmIgnoreWriteError  FileError = "NpmIgnoreWrite"
	AddScriptError       FileError = "AddScript"
)

const (
	HuskyError             ConfigError = "Husky"
	CommitLintConfigError  ConfigError = "CommitLintConfig"
	ConfigFilesCreateError ConfigError = "ConfigFilesCreate"
)

const (
	JsonUnmarshalError            ProcessError = "JsonUnmarshal"
	JsonMarshalError              ProcessError = "JsonMarshal"
	InvalidTypeForFilesToAddError ProcessError = "InvalidTypeForFilesToAdd"
	DependenciesError             ProcessError = "Dependencies"
	GitHubRepoOpenError           ProcessError = "GitHubRepoOpen"
)

const (
	AddCommitLintError           UserInputError = "AddCommitLint"
	PackageManagerSelectionError UserInputError = "PackageManagerSelection"
	ShouldPublishToNpmError      UserInputError = "ShouldPublishToNpm"
)
