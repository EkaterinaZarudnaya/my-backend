package storage

import (
	"os"
)

type Storage interface {
	UploadFile()
	GetAwsFile() *os.File
	GetFilesystemFile()
}
