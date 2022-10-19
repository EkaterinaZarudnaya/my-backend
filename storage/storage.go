package storage

import (
	"my-backend/service/file"
)

type Storage interface {
	UploadToAws(s file.CsvServise) error
	GetAwsFile() ([]byte, error)
	UploadToFilesystem(s file.CsvServise) error
	GetFilesystemFile() ([]byte, error)
}

type SaveSystem struct {
	StrMulResult [][]string
	SaveName     string
}
