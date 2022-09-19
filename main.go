package main

import (
	_ "embed"
	"my-backend/server"
)

//go:embed templates/index.html
var IndexHtml string

func main() {
	server.RouteHandler()
}
