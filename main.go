package main

import (
	_ "embed"
	"my-backend/server"
	"my-backend/server/handlers"
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
}
