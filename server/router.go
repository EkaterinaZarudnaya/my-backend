package server

import (
	"my-backend/server/handlers"
	"net/http"
)

func RouteHandler(fs handlers.FileServise) {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/upload", handlers.Upload(fs))
}
