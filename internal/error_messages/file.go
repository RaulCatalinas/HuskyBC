package error_messages

import "github.com/RaulCatalinas/HuskyBC/internal/enums"

var FILE_ERROR_MESSAGES = map[enums.FileError]string{
	enums.NotFoundError:        "The {fileName} file wasn't found in the current directory.",
	enums.ReadFileError:        generateErrorMessage("reading the {fileName} file"),
	enums.WriteFileError:       generateErrorMessage("writing to the {fileName} file"),
	enums.CreateEmptyFileError: generateErrorMessage("creating the empty file"),
	enums.CreateFolderError:    generateErrorMessage("creating the folder"),
	enums.NpmIgnoreWriteError:  generateErrorMessage("writing to the '.npmignore' file"),
	enums.AddScriptError:       generateErrorMessage("trying to add a script to the package.json file."),
}
