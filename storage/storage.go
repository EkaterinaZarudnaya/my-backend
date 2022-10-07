package storage

import (
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	s3session *s3.S3
	name      = "multiplicationResult"
	dt        = time.Now().Format("2006-01-02T15.04")
	ext       = ".csv"
	saveName  = name + dt + ext
)

const (
	REGION      = "eu-central-1"
	BUCKET_NAME = "zarudnabackendbucket"
)

type Storage interface {
	SaveFile()
}

type Filesystem struct {
	StrMulResult [][]string
	W            http.ResponseWriter
}

type Aws struct {
	StrMulResult [][]string
	W            http.ResponseWriter
}

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func NewFilesystem(strMulResult [][]string, w http.ResponseWriter) *Filesystem {
	return &Filesystem{
		StrMulResult: strMulResult,
		W:            w,
	}
}

func NewAwsSystem(strMulResult [][]string, w http.ResponseWriter) *Aws {
	return &Aws{
		StrMulResult: strMulResult,
		W:            w,
	}
}

func (fs Filesystem) SaveFile() {
	csvFile, err := os.Create(filepath.Join("saved", saveName))
	if err != nil {
		http.Error(fs.W, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()

	writeCsv(csvFile, fs.StrMulResult)
	downloadResult(csvFile, fs.W)
	fmt.Println("The file was saved successfully.")
}

func (as Aws) SaveFile() {
	tempCsvFile, err := os.CreateTemp("", saveName)
	if err != nil {
		log.Fatalln("Error creating temporary file", err)
	}
	writeCsv(tempCsvFile, as.StrMulResult)

	defer os.Remove(saveName)
	defer tempCsvFile.Close()

	tempCsvFile.Seek(0, 0)

	fmt.Println("Uploading to AWS:", saveName)

	_, err = s3session.PutObject(&s3.PutObjectInput{
		Body:   tempCsvFile,
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(saveName),
	})

	if err != nil {
		log.Fatalln("Error file uploading to AWS:", err)
	}

	downloadResult(tempCsvFile, as.W)
}

func writeCsv(file *os.File, strMulResult [][]string) {
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = ';'
	for _, row := range strMulResult {
		if err := csvWriter.Write(row); err != nil {
			log.Fatalln("Error writing record to file", err)
		}
	}
	defer csvWriter.Flush()
}

func downloadResult(file *os.File, w http.ResponseWriter) {
	tempBuffer := make([]byte, 512)
	file.Read(tempBuffer)

	FileStat, _ := file.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename="+saveName+"")
	w.Header().Set("Content-Length", FileSize)

	file.Seek(0, 0)
	io.Copy(w, file)
}
