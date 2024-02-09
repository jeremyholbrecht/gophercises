package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// open file
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		fmt.Println(rec)
	}

	// close file at end of program
	defer file.Close()
}
