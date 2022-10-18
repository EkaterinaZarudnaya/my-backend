package handlers

import (
	"fmt"
	"html/template"
	"log"
	"my-backend/service/file"
	"my-backend/service/matrix"
	"my-backend/storage"
	"my-backend/templates"
	"net/http"
	"strings"
	"time"
)

var (
	name              = "multiplicationResult"
	dt                = time.Now().Format("2006-01-02T15.04")
	ext               = ".csv"
	saveName          = name + dt + ext
	maxFileSize int64 = 600 * 1024 * 1024 //600MB
)

func Upload(w http.ResponseWriter, req *http.Request) {
	temps := templates.GetTemp()
	templateFile := template.Must(template.New("upload.html").Parse(temps["upload"]))

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
	start := time.Now()

	/*matrixA, n := matrix.ConvertItemToFlatFloat(file.Read(files[0], w))
	matrixB, m := matrix.ConvertItemToFlatFloat(file.Read(files[1], w))
	mulResult := matrix.Multiply(matrixA, matrixB, n, m)
	result := matrix.ConvertItemToString(mulResult.RawMatrix().Data, mulResult.RawMatrix().Rows, mulResult.RawMatrix().Cols)*/

	matrixA := matrix.ConvertItemToInt(file.Read(files[0], w))
	matrixB := matrix.ConvertItemToInt(file.Read(files[1], w))
	mulResult := matrix.MulMatrix(matrixA, matrixB)
	result := matrix.ConvertItemToString(mulResult)

	elapsed := time.Since(start)
	log.Printf("Time: %s", elapsed)

	//system := os.Args[1]
	system := "filesystem"
	switch strings.ToLower(system) {
	case "filesystem":
		localStorage := storage.NewFilesystem(result, w, saveName)
		localStorage.UploadFile()
		file.SaveNewCsv(file.ConvertByteToSrting(localStorage.GetFilesystemFile(), len(result)), saveName, w)
	case "aws":
		awsStorage := storage.NewAwsSystem(result, w, saveName)
		awsStorage.UploadFile()
		file.SaveNewCsv(file.ConvertByteToSrting(awsStorage.GetAwsFile(), len(result)), saveName, w)
	default:
		log.Fatalln("Invalid system parameter -", system)
		return
	}

	http.Redirect(w, req, "/upload", http.StatusSeeOther)
}
