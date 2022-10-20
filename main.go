package main

import (
	_ "embed"
	"log"
	"my-backend/server"
	"my-backend/service/file"
	"net/http"
)

func main() {
	fs := file.NewService()
	server.RouteHandler(fs)

	log.Println("Listening on localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
