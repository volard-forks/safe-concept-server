package main

import (
	"fmt"
	"log"
	"net/http"

	"safe-server/lib"
	"safe-server/server/routes"

	"github.com/joho/godotenv"
)

// serve Initializes and starts the HTTP server on the specified port.
// If errors occur logs a fatal error and terminates the application.
func serve() {
	port := lib.GetEnv("PORT")
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("launching server on %v\n", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

// loadEnvironment Loads environment variables
func loadEnvironment() {
	err := godotenv.Load("../.env")
	lib.PanicIfError(err)
}

func main() {
	loadEnvironment()
	routes.DefineRoutes()
	serve()
}
