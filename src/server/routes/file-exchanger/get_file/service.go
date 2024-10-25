package route_get_file

func getFileNameByToken(token string) string {
	// find client session by token and hash his login
	return "mock_file_name"
}

func getClientFile(token string) (string, []byte) {
	clientFileName := getFileNameByToken(token)

	return clientFileName, getFileFromFS(clientFileName)
}
