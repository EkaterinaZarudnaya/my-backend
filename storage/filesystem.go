package storage

import (
	"fmt"
	"my-backend/service/file"
	"os"
	"path/filepath"
)

func NewFilesystem(strMulResult [][]string, saveName string) *SaveSystem {
	return &SaveSystem{
		StrMulResult: strMulResult,
		SaveName:     saveName,
	}
}

func (ss SaveSystem) UploadToFilesystem(fs file.CsvServise) error {
	csvFile, err := os.Create(filepath.Join("saved", ss.SaveName))
	if err != nil {
		return err
	}
	defer csvFile.Close()
	fs.WriteCsv(csvFile, ss.StrMulResult)
	fmt.Println("The file was saved successfully into SaveSystem.")
	return nil
}

func (ss SaveSystem) GetFilesystemFile() ([]byte, error) {
	data, err := os.ReadFile("saved/" + ss.SaveName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
