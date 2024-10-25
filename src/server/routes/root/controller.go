package route_root

import (
	"io"
	"net/http"
)

func RootController(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "root")
}
