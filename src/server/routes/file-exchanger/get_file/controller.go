package route_get_file

import (
	"fmt"
	"net/http"
	"safe-server/guards"
)

func isRequestHeadersValid(r *http.Request) bool {
	// db_header, ...
	return r.Method == "GET"
}

func GetFileController(w http.ResponseWriter, r *http.Request) {
	if !isRequestHeadersValid(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientToken := r.Header.Get("token")

	if !guards.IsClientAuthorized(clientToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fileName, file := getClientFile(clientToken)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v", fileName))
	w.Write(file)

	w.WriteHeader(http.StatusOK)
}
