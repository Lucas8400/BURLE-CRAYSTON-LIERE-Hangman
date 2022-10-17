package func

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Prend un mot au hasard dans le fichier word.txt

func RandomWord() string {
	// Ouverture du fichier
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Lecture du fichier
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Récupération d'un mot au hasard
	rand.Seed(time.Now().UnixNano())
	randomWord := words[rand.Intn(len(words))]

	fmt.Println(randomWord)
	return randomWord
}
