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

	/*matrixA, n := matrix.ConvertItemToFlatFloat(file.ReadCsv(files[0], w))
	matrixB, m := matrix.ConvertItemToFlatFloat(file.ReadCsv(files[1], w))
	mulResult := matrix.Multiply(matrixA, matrixB, n, m)
	result := matrix.ConvertItemToString(mulResult.RawMatrix().Data, mulResult.RawMatrix().Rows, mulResult.RawMatrix().Cols)*/

	readResultMatrA, err := file.ReadCsv(files[0])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	matrixA := matrix.ConvertItemToInt(readResultMatrA)

	readResultMatrB, err := file.ReadCsv(files[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	matrixB := matrix.ConvertItemToInt(readResultMatrB)
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
		err := DownloadNewCsv(file.ConvertByteToSrting(localStorage.GetFilesystemFile(), len(result)), saveName, w)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
	case "aws":
		awsStorage := storage.NewAwsSystem(result, w, saveName)
		awsStorage.UploadFile()
		err := DownloadNewCsv(file.ConvertByteToSrting(awsStorage.GetAwsFile(), len(result)), saveName, w)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
	default:
		log.Fatalln("Invalid system parameter -", system)
		return
	}

	http.Redirect(w, req, "/upload", http.StatusSeeOther)
}
