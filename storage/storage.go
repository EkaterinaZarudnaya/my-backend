package storage

import (
	"my-backend/server/handlers"
	"net/http"
)

type Storage interface {
	UploadToAws(s handlers.FileServise)
	GetAwsFile() []byte
	UploadFile(s handlers.FileServise)
	GetFilesystemFile() []byte
}

type SaveSystem struct {
	StrMulResult [][]string
	W            http.ResponseWriter
	SaveName     string
}
