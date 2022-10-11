package storage

import (
	"fmt"
	"log"
	"my-backend/service"
	"net/http"
	"os"
	"path/filepath"
)

type Filesystem struct {
	StrMulResult [][]string
	W            http.ResponseWriter
	SaveName     string
}

func NewFilesystem(strMulResult [][]string, w http.ResponseWriter, saveName string) *Filesystem {
	return &Filesystem{
		StrMulResult: strMulResult,
		W:            w,
		SaveName:     saveName,
	}
}

func (fs Filesystem) UploadFile() {
	csvFile, err := os.Create(filepath.Join("saved", fs.SaveName))
	if err != nil {
		http.Error(fs.W, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()
	service.WriteCsv(csvFile, fs.StrMulResult)
	fmt.Println("The file was saved successfully into Filesystem.")
}

func (fs Filesystem) GetFilesystemFile() []byte {
	data, err := os.ReadFile("saved/" + fs.SaveName)
	if err != nil {
		log.Fatalln("Error file getting from filesystem:", err)
	}
	return data
}
