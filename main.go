package main

import (
	"encoding/csv"
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
	score := 0

	records := ReadCSV("problems.csv")

	for i, strings := range records {
		fmt.Printf("Problem #%d: %v = ", i+1, strings[0])

		var guess string
		fmt.Scanln(&guess)

		if guess == strings[1] {
			score += 1
		}

	}

	fmt.Printf("Your scored: %v out of %v\n", score, len(records))
}
