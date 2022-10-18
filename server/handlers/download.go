package handlers

import (
	"fmt"
	"io"
	"log"
	"my-backend/service/file"
	"net/http"
	"os"
	"strconv"
)

func handleDownload(file *os.File, w http.ResponseWriter, saveName string) {
	FileStat, _ := file.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename="+saveName+"")
	w.Header().Set("Content-Length", FileSize)

	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Printf("Error setting offset: %v\n", err)
		return
	}
	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Printf("Copy error: %v\n", err)
		return
	}
}

func DownloadNewCsv(newStrResult [][]string, saveName string, w http.ResponseWriter) error {
	newCsvFile, err := os.CreateTemp("", saveName)
	if err != nil {
		log.Fatalln("Error creating temporary file", err)
		return err
	}

	err = file.WriteCsv(newCsvFile, newStrResult)
	if err != nil {
		return err
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatalln("Error removing file: ", err)
		}
	}(saveName)

	defer func(newCsvFile *os.File) {
		err := newCsvFile.Close()
		if err != nil {
			log.Fatalln("Error closing file: ", err)
		}
	}(newCsvFile)

	_, err = newCsvFile.Seek(0, 0)
	if err != nil {
		return err
	}

	handleDownload(newCsvFile, w, saveName)
	return nil
}
