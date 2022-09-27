package service

import (
	"encoding/csv"
	"fmt"
	"github.com/dimchansky/utfbom"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func ReadFile(file *multipart.FileHeader, w http.ResponseWriter) [][]string {
	f, err := file.Open()
	fl, _ := utfbom.Skip(f)
	if err != nil {
		fmt.Printf("File retrieval error: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return nil
	}
	csvReader := csv.NewReader(fl)
	csvReader.FieldsPerRecord = -1
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Can't read the file: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return nil
	}
	f.Close()
	return data
}

func SaveFile(strMulResult [][]string, w http.ResponseWriter) {
	name := "multiplicationResult"
	dt := time.Now().Format("2006-01-02T15.04")
	ext := ".csv"

	saveName := name + dt + ext
	filePath := filepath.Join("saved", saveName)

	csvFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = ';'
	for _, row := range strMulResult {
		if err := csvWriter.Write(row); err != nil {
			log.Fatalln("Error writing record to file", err)
		}
	}
	defer csvWriter.Flush()

	fmt.Println("The file was saved successfully.")
}
