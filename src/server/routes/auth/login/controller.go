package route_login

import (
	"encoding/json"
	"net/http"
	route_login "safe-server/services/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO link with db
	if !isValidCredentials(req.Username, req.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := route_login.GenerateToken(req.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func isValidCredentials(username, password string) bool {
	// TODO
	return true
}
