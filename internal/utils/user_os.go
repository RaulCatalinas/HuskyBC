package utils

func CreateFolderOrFileIfNotExists(path string, isFolder bool) {
	exists := CheckIfExists(path)

	if !exists {
		createFolderOrFile(path, isFolder)
	}
}
