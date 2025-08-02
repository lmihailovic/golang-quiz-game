package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func ReadCSV(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records
}

func main() {
	problemsFile := flag.String("file", "problems.csv", "path to problems file")
	flag.Parse()

	records := ReadCSV(*problemsFile)
	score := 0

	for i, strings := range records {
		fmt.Printf("Problem #%d: %v = ", i+1, strings[0])

		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			log.Fatal(err)
		}

		if answer == strings[1] {
			score += 1
		}
	}

	fmt.Printf("Your scored: %v out of %v\n", score, len(records))
}
