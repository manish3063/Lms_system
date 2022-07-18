package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// var records = readCsvFile("./books.csv")
func main() {
	createDBConnection()

	//records := readCsvFile("./books.csv")
	//importCSV(records)

	r := gin.Default()
	setupRoutes(r)
	r.Run() //

}

func importCSV(records [][]string) {

	for i := 1; i < 1000; i++ {

		sqlStatement := `INSERT INTO books(book_id,book_name,book_author,book_image)
    VALUES ($1, $2, $3, $4)`

		_, err2 := DB.Exec(sqlStatement, records[i][0], records[i][1], records[i][2], records[i][5])

		if err2 != nil {
			log.Println("ERror in insert: ", err2)
		}

	}
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	// csvReader := csv.NewReader(f)
	// records, err := csvReader.ReadAll()

	reader := csv.NewReader(bufio.NewReader(f))
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
