package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word     string
	ToFind   string
	Attempts int
}

func Scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func main() {
	var hangman HangManData
	hangman.Init()
	for hangman.Attempts > 0 {
		fmt.Println("Mot à trouver:", hangman.Word)
		fmt.Println("Nombre d'essais restants:", hangman.Attempts)
		fmt.Println("Entrez une lettre:")
		input := Scan()
		if VerifyLetter(input, hangman.ToFind) {
			hangman.Attempts--
		}
		var indexes []int
		for index, letter := range hangman.ToFind {
			if input == string(letter) {
				if VerifyIndex(indexes, index) {
					indexes = append(indexes, index)
				}
			}
		}
		for _, index := range indexes {
			hangman.Word = Replace(hangman.Word, input, index)
		}
		if hangman.Word == hangman.ToFind {
			fmt.Println("Bravo, vous avez trouvé le mot:", hangman.Word)
			break
		}
	}
	if hangman.Attempts == 0 {
		fmt.Println("Vous avez perdu, le mot était:", hangman.ToFind)
	}
}

func Replace(word string, input string, index int) string {
	var new_word string
	for i, letter := range word {
		if i == index {
			new_word += input
		} else {
			new_word += string(letter)
		}
	}
	return new_word
}

func (h *HangManData) Init() {
	h.ToFind = RandomWord()
	for range h.ToFind {
		h.Word += "_"
	}
	h.Attempts = 10
}

func VerifyLetter(input string, word string) bool {
	for _, letter := range word {
		if input == string(letter) {
			return false
		}
	}
	return true
}

func VerifyIndex(tab []int, index int) bool {
	for _, element := range tab {
		if index == element {
			return false
		}
	}
	return true
}

func RandomWord() string {
	readFile, _ := os.Open("words.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	rand.Seed(time.Now().Unix())
	mot := lines[rand.Intn(len(lines))]
	return mot
}
