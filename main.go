package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a file in format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit to answer in second")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Couldn't open csv file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		fmt.Println("Failed to parse csv file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	<-timer.C

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {

			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d of %d", correct, len(problems))
			return
		case answer := <-answerCh:
			// default:
			// 	fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			// 	var answer string
			// 	fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				correct++
			}

		}
	}
	fmt.Printf("Score: %d out of %d", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, lines := range lines {
		ret[i] = problem{
			q: lines[0],
			a: strings.TrimSpace(lines[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
