package main

import (
	_ "embed"
	"my-backend/server"
	"my-backend/server/handlers"
)

//go:embed templates/index.html
var indexHtml string

func main() {
	handlers.IndexHtml = indexHtml
	server.RouteHandler()
}
