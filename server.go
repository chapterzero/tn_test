package main

import (
	"github.com/chapterzero/tn_test/router"
	"github.com/chapterzero/tn_test/server"
	"log"
	"net/http"
	"path/filepath"
)

var appConfig server.AppConfig

func init() {
	configAbsPath, err := filepath.Abs("./app.config.json")
	if err != nil {
		panic("Could not find app.config.json")
	}

	// loading config file and initializing service class
	appConfig = server.LoadConfigFromFile(configAbsPath)
	server.SetDbConfig(appConfig.Db)
}

func main() {
	log.Println("HTTP Server will be running at http://localhost:8777")
	err := http.ListenAndServe(":8777", router.CreateRouter())

	if err != nil {
		panic(err)
	}
}
