package handlers

import (
	"fmt"
	"html/template"
	"log"
	"my-backend/service"
	"my-backend/storage"
	"net/http"
	"strings"
	"time"
)

var (
	UploadHtml  string
	System      string
	name              = "multiplicationResult"
	dt                = time.Now().Format("2006-01-02T15.04")
	ext               = ".csv"
	saveName          = name + dt + ext
	maxFileSize int64 = 5 * 1024 * 1024 //5MB
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
	req.Body = http.MaxBytesReader(w, req.Body, maxFileSize)

	err := req.ParseMultipartForm(maxFileSize)
	if err != nil {
		fmt.Printf("Parse error - %v\n", err.Error())
		http.Error(w, "Request Too Large", http.StatusRequestEntityTooLarge)
		return
	}

	files := req.MultipartForm.File["files"]
	fmt.Println("Retrieving the files successfully.")

	matrixA, n := service.ConvertItemToFlatFloat(service.ReadFile(files[0], w))
	matrixB, m := service.ConvertItemToFlatFloat(service.ReadFile(files[1], w))
	mulResult := service.MulMatrix(matrixA, matrixB, n, m)
	result := service.ConvertItemToString(mulResult.RawMatrix().Data, mulResult.RawMatrix().Rows, mulResult.RawMatrix().Cols)

	switch strings.ToLower(System) {
	case "filesystem":
		localStorage := storage.NewFilesystem(result, w, saveName)
		localStorage.UploadFile()
		service.SaveNewCsv(service.ConvertByteToSrting(localStorage.GetFilesystemFile(), len(result)), saveName, w)
	case "aws":
		awsStorage := storage.NewAwsSystem(result, w, saveName)
		awsStorage.UploadFile()
		service.SaveNewCsv(service.ConvertByteToSrting(awsStorage.GetAwsFile(), len(result)), saveName, w)
	default:
		log.Fatalln("Invalid system parameter -", System)
		return
	}

	http.Redirect(w, req, "/upload", http.StatusSeeOther)
}
