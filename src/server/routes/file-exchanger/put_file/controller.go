package route_put_file

import (
	"net/http"
)

func isRequestHeadersValid() bool {
	// db_header, method, ...
	return true
}

func isFileValid() bool {
	// size, format, content
	return true
}

func getFileNameFromRequest(r *http.Request) string {
	return "mock_file_name"
}

func getFileContentFromRequest(r *http.Request) string {
	return "mock_file_content"
}

func PutFileController(w http.ResponseWriter, r *http.Request) {
	if !isRequestHeadersValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !isFileValid() {
		// w.WriteHeader(http.StatusRequestEntityTooLarge)
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	saveFile(getFileNameFromRequest(r), getFileContentFromRequest(r))

	w.WriteHeader(http.StatusOK)
}
