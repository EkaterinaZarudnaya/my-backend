package storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"my-backend/service"
	"net/http"
	"os"
)

var (
	s3session *s3.S3
)

const (
	REGION      = "eu-central-1"
	BUCKET_NAME = "zarudnabackendbucket"
)

type Aws struct {
	StrMulResult [][]string
	W            http.ResponseWriter
	SaveName     string
}

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func NewAwsSystem(strMulResult [][]string, w http.ResponseWriter, saveName string) *Aws {
	return &Aws{
		StrMulResult: strMulResult,
		W:            w,
		SaveName:     saveName,
	}
}

func (a Aws) UploadFile() {
	tempCsvFile, err := os.CreateTemp("", a.SaveName)
	if err != nil {
		log.Fatalln("Error creating temporary file", err)
	}
	service.WriteCsv(tempCsvFile, a.StrMulResult)

	defer os.Remove(a.SaveName)
	defer tempCsvFile.Close()

	tempCsvFile.Seek(0, 0)

	fmt.Println("Uploading to AWS:", a.SaveName)

	_, err = s3session.PutObject(&s3.PutObjectInput{
		Body:   tempCsvFile,
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(a.SaveName),
	})
	if err != nil {
		log.Fatalln("Error file uploading to AWS:", err)
	}
}

func (a Aws) GetAwsFile() []byte {
	fmt.Println("Downloading: ", a.SaveName)
	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(a.SaveName),
	})
	if err != nil {
		log.Fatalln("Error file getting from AWS:", err)
	}
	body, err := io.ReadAll(resp.Body)
	return body
}
