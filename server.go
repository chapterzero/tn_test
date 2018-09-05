package main

import (
	"github.com/chapterzero/tn_test/server"
	"log"
	"net/http"
	"path/filepath"
)

func init() {
	configAbsPath, err := filepath.Abs("./app.config.json")
	if err != nil {
		panic("Could not find app.config.json")
	}

	// loading config file and initializing service class
	appConfig := server.LoadConfigFromFile(configAbsPath)
	log.Println(appConfig)
}

func main() {
	log.Println("Server will be running at port 8777")
	http.ListenAndServe(":8777", server.CreateRouter())
}
