package server

import (
	"my-backend/server/handlers"
	"my-backend/service/file"
	"my-backend/service/mongodb"
	"net/http"
)

func RouteHandler(fs file.CsvServise, ms mongodb.ResultDownloadsServise) {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/upload", handlers.Upload(fs, ms))
}
