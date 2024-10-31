package guards

import "safe-server/services/auth"

func IsClientAuthorized(token string) bool {
	if token == "" {
		return false
	}

	userID, err := auth.ValidateToken(token)
	if err != nil {
		return false
	}

	return userID != ""
}
