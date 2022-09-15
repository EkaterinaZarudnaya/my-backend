package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FileHandler(Filename string, file multipart.File, w http.ResponseWriter) {
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
