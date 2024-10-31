package route_get_file

import (
	"os"
	"safe-server/lib"
)

// getFileFromFS retrieves a file from the filesystem using the configured FS_FOLDER_PATH
// environment variable as the base directory.
//
// Parameters:
//   - fileName: The name of the file
//
// Returns:
//   - []byte: The contents of the file
func getFileFromFS(fileName string) []byte {
	fsFolderPath := lib.GetEnv("FS_FOLDER_PATH")

	file, err := os.ReadFile(fsFolderPath + fileName)

	lib.PanicIfError(err)

	return file
}
