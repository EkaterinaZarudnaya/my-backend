package handlers

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var IndexHtml string

func Upload(w http.ResponseWriter, req *http.Request) {
	templateFile := template.Must(template.New("index.html").Parse(IndexHtml))

	if req.Method == http.MethodPost {
		handleUpload(w, req)
		return
	}

	templateFile.ExecuteTemplate(w, "index.html", nil)
}

func handleUpload(w http.ResponseWriter, req *http.Request) {
	var maxFileSize int64 = 5 * 1024 * 1024 //5MB

	req.Body = http.MaxBytesReader(w, req.Body, maxFileSize)

	err := req.ParseMultipartForm(maxFileSize)
	if err != nil {
		fmt.Printf("Parse error - %v\n", err.Error())
		http.Redirect(w, req, "/?success=false", http.StatusRequestEntityTooLarge)
		return
	}

	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		fmt.Printf("File retrieval error: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println("Retrieving the file successfully.")
	defer file.Close()

	saveFile(fileHeader.Filename, file, w)

	http.Redirect(w, req, "/?success=true", http.StatusSeeOther)
}

func saveFile(Filename string, file multipart.File, w http.ResponseWriter) {
	dt := time.Now().Format("2006-01-02T15.04")
	name := strings.TrimSuffix(Filename, filepath.Ext(Filename))
	ext := filepath.Ext(Filename)

	saveName := name + dt + ext
	filePath := filepath.Join("uploads", saveName)

	dest, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("The file was saved successfully.")
}
