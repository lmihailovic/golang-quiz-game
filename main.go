package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	fmt.Printf("\nYou scored: %v out of %v\n", score, totalProblems)
	os.Exit(0)
}

func main() {
	problemsFile := flag.String("f", "problems.csv", "path to problems file")
	gameTime := flag.Int("t", 30, "time to play in seconds")
	flag.Parse()

	records := ReadCSV(*problemsFile)
	score := 0

	fmt.Println("Press any key to start.")
	fmt.Scanln() // no need to check for errors

	time.AfterFunc(time.Duration(*gameTime)*time.Second, func() {
		EndGame(score, len(records))
	})

	for i, problem := range records {
		fmt.Printf("Problem #%d: %v = ", i+1, problem[0])

		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			log.Fatal(err)
		}

		answer = strings.Trim(answer, " ")
		answer = strings.ToLower(answer)

		if answer == problem[1] {
			score += 1
		}
	}

	EndGame(score, len(records))
}
