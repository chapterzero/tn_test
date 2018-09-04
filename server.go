package main

import (
	"log"
	"net/http"
	"github.com/chapterzero/tn_test/router"
)

type config {
}

func main() {
	log.Println("Server will be running at port 8777")
	http.ListenAndServe(":8777", router.CreateRouter())
}
