package main

import (
	_ "embed"
	"log"
	"my-backend/server"
	"my-backend/service/file"
	"my-backend/service/mongodb"
	"net/http"
)

func main() {
	fs := file.NewService()
	ms := mongodb.NewService()

	server.RouteHandler(fs, ms)

	log.Println("Listening on localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
