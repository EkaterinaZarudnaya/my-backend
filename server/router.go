package server

import (
	"my-backend/server/handlers"
	"my-backend/service/file"
	"net/http"
)

func RouteHandler(fs file.CsvServise) {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/upload", handlers.Upload(fs))
}
