package handlers

import (
	"fmt"
	"io"
	"log"
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

func DownloadNewCsv(fs FileServise, newStrResult [][]string, saveName string, w http.ResponseWriter) error {
	newCsvFile, err := os.CreateTemp("", saveName)
	if err != nil {
		log.Fatalln("Error creating temporary file", err)
		return err
	}

	err = fs.WriteCsv(newCsvFile, newStrResult)
	if err != nil {
		return err
	}

	defer os.Remove(saveName)
	defer newCsvFile.Close()

	_, err = newCsvFile.Seek(0, 0)
	if err != nil {
		return err
	}

	handleDownload(newCsvFile, w, saveName)

	return nil
}
