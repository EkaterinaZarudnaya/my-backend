package main

import (
	_ "embed"
	"my-backend/server"
	"my-backend/server/handlers"
)

//go:embed templates/index.html
var indexHtml string

//go:embed templates/upload.html
var uploadHtml string

func main() {
	handlers.IndexHtml = indexHtml
	handlers.UploadHtml = uploadHtml
	server.RouteHandler()
}
