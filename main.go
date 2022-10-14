package main

import (
	_ "embed"
	"log"
	"my-backend/server"
	"net/http"
)

func main() {
	server.RouteHandler()
	log.Println("Listening on localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
