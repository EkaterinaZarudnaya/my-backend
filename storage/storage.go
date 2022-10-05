package storage

import (
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func NewAwsSystem(strMulResult [][]string) *Aws {
	return &Aws{
		StrMulResult: strMulResult,
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
