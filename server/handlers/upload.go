package handlers

import (
	"fmt"
	"html/template"
	"log"
	"my-backend/service/file"
	"my-backend/service/matrix"
	"my-backend/service/mongodb"
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

func Upload(fs file.CsvServise, ms mongodb.ResultDownloadsServise) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			start := time.Now()
			handleUpload(fs, ms)(w, req)
			elapsed := time.Since(start)
			log.Printf("Time: %s", elapsed)
		case http.MethodGet:
			temps := templates.GetTemp()
			templateFile := template.Must(template.New("upload.html").Parse(temps["upload"]))
			templateFile.ExecuteTemplate(w, "upload.html", nil)
		}
	}
}

func handleUpload(fs file.CsvServise, ms mongodb.ResultDownloadsServise) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.Body = http.MaxBytesReader(w, req.Body, maxFileSize)

		err := req.ParseMultipartForm(maxFileSize)
		if err != nil {
			fmt.Printf("Parse error - %v\n", err.Error())
			http.Error(w, "Request Too Large", http.StatusRequestEntityTooLarge)
			return
		}

		files := req.MultipartForm.File["files"]
		fmt.Println("Retrieving the files successfully.")

		readResultMatrA, err := fs.ReadCsv(files[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		matrixA := matrix.MatString(readResultMatrA).ToInt()

		readResultMatrB, err := fs.ReadCsv(files[1])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		matrixB := matrix.MatString(readResultMatrB).ToInt()

		mulResult := matrix.MulMatrix(matrixA, matrixB)
		result := mulResult.ToString()

		//system := os.Args[1]
		system := "filesystem"
		switch strings.ToLower(system) {
		case "filesystem":
			localStorage := storage.NewFilesystem(result, saveName)
			err := localStorage.UploadToFilesystem(fs)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			receivedFile, err := localStorage.GetFilesystemFile()
			if err != nil {
				log.Fatalln("Error file getting from filesystem:", err)
			}

			err = DownloadNewCsv(fs, fs.ConvertByteToSrting(receivedFile, len(result)), saveName, w)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				http.Error(w, "Bad Request", http.StatusBadRequest)
			}

		case "aws":
			awsStorage := storage.NewAwsSystem(result, saveName)
			err := awsStorage.UploadToFilesystem(fs)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			receivedFile, err := awsStorage.GetAwsFile()
			if err != nil {
				log.Fatalln("Error file getting from AWS:", err)
			}

			err = DownloadNewCsv(fs, fs.ConvertByteToSrting(receivedFile, len(result)), saveName, w)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				http.Error(w, "Bad Request", http.StatusBadRequest)
			}
		default:
			log.Fatalln("Invalid system parameter -", system)
			return
		}

		inputData := map[string]string{
			"system":    strings.ToLower(system),
			"fileName":  saveName,
			"createdAt": dt,
		}

		err = ms.InitConnections()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Redirect(w, req, "/upload", http.StatusSeeOther)
		}

		ms.NewData(inputData)

		err = ms.InsertOneIntoCollect()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Redirect(w, req, "/upload", http.StatusSeeOther)
		}

		ms.Disconnections()

		http.Redirect(w, req, "/upload", http.StatusSeeOther)
	}
}
