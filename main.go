package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "problem.csv", "a file in format of 'question,answer'")
	flag.Parse()
	_ = csvFilename
}
