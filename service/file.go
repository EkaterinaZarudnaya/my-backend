package service

import (
	"encoding/csv"
	"fmt"
	"github.com/dimchansky/utfbom"
	"mime/multipart"
	"net/http"
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
