package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
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

func EndGame(score int, totalProblems int) {
	fmt.Printf("\nYour scored: %v out of %v\n", score, totalProblems)
	os.Exit(0)
}

func main() {
	problemsFile := flag.String("f", "problems.csv", "path to problems file")
	gameTime := flag.Int("t", 30, "time to play in seconds")
	flag.Parse()

	records := ReadCSV(*problemsFile)
	score := 0

	fmt.Println("Press any key to start.")
	fmt.Scanln()

	time.AfterFunc(time.Duration(*gameTime)*time.Second, func() {
		EndGame(score, len(records))
	})

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

	EndGame(score, len(records))
}
