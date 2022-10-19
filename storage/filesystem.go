package storage

import (
	"fmt"
	"log"
	"my-backend/server/handlers"
	"net/http"
	"os"
	"path/filepath"
)

func NewFilesystem(strMulResult [][]string, w http.ResponseWriter, saveName string) *SaveSystem {
	return &SaveSystem{
		StrMulResult: strMulResult,
		W:            w,
		SaveName:     saveName,
	}
}

func (ss SaveSystem) UploadFile(fs handlers.FileServise) {
	csvFile, err := os.Create(filepath.Join("saved", ss.SaveName))
	if err != nil {
		http.Error(ss.W, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()
	fs.WriteCsv(csvFile, ss.StrMulResult)
	fmt.Println("The file was saved successfully into SaveSystem.")
}

func (ss SaveSystem) GetFilesystemFile() []byte {
	data, err := os.ReadFile("saved/" + ss.SaveName)
	if err != nil {
		log.Fatalln("Error file getting from filesystem:", err)
	}
	return data
}
