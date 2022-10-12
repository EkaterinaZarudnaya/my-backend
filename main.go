package main

import (
	_ "embed"
	"log"
	"my-backend/server"
	"my-backend/server/handlers"
	"net/http"
	"os"
)

var (
	//go:embed templates/index.html
	indexHtml string
	//go:embed templates/upload.html
	uploadHtml string
)

func main() {
	handlers.IndexHtml = indexHtml
	handlers.UploadHtml = uploadHtml
	handlers.System = os.Args[1]
	server.RouteHandler()
	
	log.Println("Listening on localhost:8080 ...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
