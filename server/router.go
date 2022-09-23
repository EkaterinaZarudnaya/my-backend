package server

import (
	"log"
	"my-backend/server/handlers"
	"net/http"
)

func RouteHandler() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/upload", handlers.Upload)

	log.Println("Listening on localhost:8080 ...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
