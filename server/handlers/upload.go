package handlers

import (
	"fmt"
	"html/template"
	"log"
	"my-backend/service"
	"my-backend/storage"
	"net/http"
	"strings"
)

var (
	UploadHtml string
	System     string
)

func Upload(w http.ResponseWriter, req *http.Request) {
	templateFile := template.Must(template.New("upload.html").Parse(UploadHtml))

	if req.Method == http.MethodPost {
		handleUpload(w, req)
		return
	}

	templateFile.ExecuteTemplate(w, "upload.html", nil)
}

func handleUpload(w http.ResponseWriter, req *http.Request) {
	var maxFileSize int64 = 5 * 1024 * 1024 //5MB
	req.Body = http.MaxBytesReader(w, req.Body, maxFileSize)

	err := req.ParseMultipartForm(maxFileSize)
	if err != nil {
		fmt.Printf("Parse error - %v\n", err.Error())
		http.Error(w, "Request Too Large", http.StatusRequestEntityTooLarge)
		return
	}

	files := req.MultipartForm.File["files"]
	fmt.Println("Retrieving the files successfully.")
	matrixA := service.ConvertItemToInt(service.ReadFile(files[0], w))
	matrixB := service.ConvertItemToInt(service.ReadFile(files[1], w))
	result := service.ConvertItemToString(service.MulMatrix(matrixA, matrixB))

	switch strings.ToLower(System) {
	case "filesystem":
		storage.NewFilesystem(result, w).SaveFile()
	case "aws":
		storage.NewAwsSystem(result).SaveFile()
	default:
		log.Fatalln("Invalid system parameter -", System)
		return
	}

	http.Redirect(w, req, "/upload", http.StatusSeeOther)
}
