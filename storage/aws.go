package storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"my-backend/server/handlers"
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

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func NewAwsSystem(strMulResult [][]string, w http.ResponseWriter, saveName string) *SaveSystem {
	return &SaveSystem{
		StrMulResult: strMulResult,
		W:            w,
		SaveName:     saveName,
	}
}

func (ss SaveSystem) UploadToAws(fs handlers.FileServise) {
	tempCsvFile, err := os.CreateTemp("", ss.SaveName)
	if err != nil {
		log.Fatalln("Error creating temporary file", err)
	}

	fs.WriteCsv(tempCsvFile, ss.StrMulResult)

	defer os.Remove(ss.SaveName)
	defer tempCsvFile.Close()

	tempCsvFile.Seek(0, 0)

	fmt.Println("Uploading to AWS:", ss.SaveName)

	_, err = s3session.PutObject(&s3.PutObjectInput{
		Body:   tempCsvFile,
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(ss.SaveName),
	})
	if err != nil {
		log.Fatalln("Error file uploading to AWS:", err)
	}
}

func (ss SaveSystem) GetAwsFile() []byte {
	fmt.Println("Downloading: ", ss.SaveName)
	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(ss.SaveName),
	})
	if err != nil {
		log.Fatalln("Error file getting from AWS:", err)
	}
	body, err := io.ReadAll(resp.Body)
	return body
}
