package file

import (
	"encoding/csv"
	"github.com/dimchansky/utfbom"
	"mime/multipart"
	"os"
	"strings"
)

func ReadCsv(file *multipart.FileHeader) ([][]string, error) {
	f, err := file.Open()
	fl, _ := utfbom.Skip(f)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(fl)
	csvReader.FieldsPerRecord = -1
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	f.Close()
	return data, nil
}

func WriteCsv(file *os.File, strMulResult [][]string) error {
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = ';'
	for _, row := range strMulResult {
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	defer csvWriter.Flush()
	return nil
}

func ConvertByteToSrting(body []byte, n int) [][]string {
	strBody := string(body)
	cont := strings.FieldsFunc(strBody, Split)

	strResultBody := make([][]string, n)
	for i := 0; i < n; i++ {
		strResultBody[i] = make([]string, n)
		for j := 0; j < n; j++ {
			strResultBody[i][j] = cont[i*n+j]
		}
	}
	return strResultBody
}

func Split(r rune) bool {
	return r == ';' || r == '\n'
}
