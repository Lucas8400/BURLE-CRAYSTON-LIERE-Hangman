package main

import (
	"bufio"
	"os"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
}

func main() {
	readFile, _ := os.Open("words.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
}
