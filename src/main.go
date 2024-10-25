package main

import (
	"fmt"
	"log"
	"net/http"

	"safe-server/lib"
	"safe-server/server/routes"

	"github.com/joho/godotenv"
)

func serve() {
	port := lib.GetEnv("PORT")
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("launching server on %v\n", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func loadEnvironment() {
	err := godotenv.Load("../.env")
	lib.PanicIfError(err)
}

func main() {
	loadEnvironment()
	routes.DefineRoutes()
	serve()
}
