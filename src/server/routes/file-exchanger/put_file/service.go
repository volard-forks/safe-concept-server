package route_put_file

func getFileNameToSave(fileName string) string {
	// get uuid or hash from fileName
	return fileName
}

func getSanitizedFileContent(fileContent string) string {
	// sanitize content with libs from js, css and html injections
	return fileContent
}

func saveFile(fileName string, fileContent string) {
	fileNameToSave := getFileNameToSave(fileName)

	fileContentToSave := getSanitizedFileContent(fileContent)

	writeFileOnFS(fileNameToSave, fileContentToSave)
}