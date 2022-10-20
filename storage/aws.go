package storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"my-backend/service/file"
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

func NewAwsSystem(strMulResult [][]string, saveName string) *SaveSystem {
	return &SaveSystem{
		StrMulResult: strMulResult,
		SaveName:     saveName,
	}
}

func (ss SaveSystem) UploadToAws(fs file.CsvServise) error {
	tempCsvFile, err := os.CreateTemp("", ss.SaveName)
	if err != nil {
		return err
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
		return err
	}
	return nil
}

func (ss SaveSystem) GetAwsFile() ([]byte, error) {
	fmt.Println("Downloading: ", ss.SaveName)
	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(ss.SaveName),
	})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	return body, nil
}
