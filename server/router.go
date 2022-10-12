package server

import (
	"my-backend/server/handlers"
	"net/http"
)

func RouteHandler() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/upload", handlers.Upload)
}
