package routes

import (
	"net/http"

	route_get_file "safe-server/server/routes/file-exchanger/get_file"
	route_put_file "safe-server/server/routes/file-exchanger/put_file"
	route_root "safe-server/server/routes/root"
)

func DefineRoutes() {
	http.HandleFunc("/", route_root.RootController)
	http.HandleFunc("/get_file", route_get_file.GetFileController)
	http.HandleFunc("/put_file", route_put_file.PutFileController)
}
