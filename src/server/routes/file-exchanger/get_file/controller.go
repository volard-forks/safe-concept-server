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

// GetFileController Handles HTTP GET requests for file downloads. It validates the request,
// authenticates the client via token, and serves the requested file.
//
// Parameters:
//   - w: HTTP response writer for sending the response
//   - r: HTTP request containing headers and token
//
// The function performs the following steps:
//  1. Validates request headers
//  2. Authenticates client token
//  3. Retrieves and serves the file
//
// HTTP Status Codes:
//   - 200 OK: File successfully retrieved and sent
//   - 400 Bad Request: Invalid request headers
//   - 401 Unauthorized: Invalid or missing token
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
