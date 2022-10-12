package file

import (
	"encoding/csv"
	"fmt"
	"github.com/dimchansky/utfbom"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Read(file *multipart.FileHeader, w http.ResponseWriter) [][]string {
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

func WriteCsv(file *os.File, strMulResult [][]string) {
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = ';'
	for _, row := range strMulResult {
		if err := csvWriter.Write(row); err != nil {
			log.Fatalln("Error writing record to file", err)
		}
	}
	defer csvWriter.Flush()
}

func ConvertByteToSrting(body []byte, n int) [][]string {
	strBoby := string(body)
	cont := strings.FieldsFunc(strBoby, Split)

	strResultBody := make([][]string, n)
	for i := 0; i < n; i++ {
		strResultBody[i] = make([]string, n)
		for j := 0; j < n; j++ {
			strResultBody[i][j] = cont[i*n+j]
		}
	}
	return strResultBody
}

func Split(r rune) bool {
	return r == ';' || r == '\n'
}

func Download(file *os.File, w http.ResponseWriter, saveName string) {
	FileStat, _ := file.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename="+saveName+"")
	w.Header().Set("Content-Length", FileSize)

	file.Seek(0, 0)
	io.Copy(w, file)
}

func SaveNewCsv(newStrResult [][]string, saveName string, w http.ResponseWriter) {
	newCsvFile, err := os.CreateTemp("", saveName)
	if err != nil {
		log.Fatalln("Error creating temporary file", err)
	}
	WriteCsv(newCsvFile, newStrResult)
	defer os.Remove(saveName)
	defer newCsvFile.Close()
	newCsvFile.Seek(0, 0)
	Download(newCsvFile, w, saveName)
}
