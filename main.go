package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problem.csv", "a file in format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Couldn't open csv file: %s\n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		fmt.Println("Failed to parse csv file.")
	}
	problems := parseLines(lines)
	fmt.Println(problems)

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
			fmt.Println("correct")
		}
	}
	fmt.Printf("Score: %d out of %d", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	give := make([]problem, len(lines))

	for i, lines := range lines {
		give[i] = problem{
			q: lines[0],
			a: lines[1],
		}
	}
	return give
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
