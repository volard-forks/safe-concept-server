package route_get_file

import (
	"os"
	"safe-server/lib"
)

func getFileFromFS(fileName string) []byte {
	fsFolderPath := lib.GetEnv("FS_FOLDER_PATH")

	file, err := os.ReadFile(fsFolderPath + fileName)

	lib.PanicIfError(err)

	return file
}
