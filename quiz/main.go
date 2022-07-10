package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	// parse user flags
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	// read hte csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// get the lines from the file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the CSV file: %s\n", *csvFilename))
	}

	// instantiate variables
	problems := parseLines(lines)
	counter := 0

func parseLines(lines [][]string) []problem {
	// use exact length, to reduce the overhead of the append function since we know the exact length.
	ret := make([]problem, len(lines))

	// iterate over the lines and assign the values to the problem struct
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
