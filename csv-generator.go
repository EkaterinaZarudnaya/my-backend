package main

/*
import (
	"encoding/csv"
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var size int
	flag.IntVar(&size, "size", 3, "specify the dimension of the square matrix")
	var filename string
	flag.StringVar(&filename, "filename", "matrix", "specify the filename")
	flag.Parse()

	randMatrix := make([][]string, size)
	for i := 0; i < size; i++ {
		randMatrix[i] = make([]string, size)
	}
	generateMatrix(randMatrix)

	newCsw, err := os.Create(filename + ".csv")
	defer newCsw.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(newCsw)
	defer w.Flush()

	err = writeCsv(newCsw, randMatrix)
	if err != nil {
		log.Fatalln("Error writing record to file", err)
		return
	}
}

func generateMatrix(randMatrix [][]string) {
	for i, innerArray := range randMatrix {
		for j := range innerArray {
			item := rand.Intn(100)
			randMatrix[i][j] = strconv.Itoa(item)
		}
	}
}

func writeCsv(file *os.File, strMulResult [][]string) error {
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
*/
