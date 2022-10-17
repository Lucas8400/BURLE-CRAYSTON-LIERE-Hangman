package classic

import (
	"fmt"
)

type HangManData struct {
	Word             string // Word composed of '_', ex: H_ll_
	ToFind           string // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func (h *HangManData) Init(){
	for index,letter range := ToFind{
		if index 
	}
	h.ToFind		   = RandomWord()
	h.Attempts 		   = int
	h.HangmanPositions = [10]string
}

func (h *HangManData) verify(){
	var lettre_choisi string
	fmt.Println("Choose :")
	fmt.Scan(lettre_choisi)
	if h.Attempts > 0{
		for index,letter := range h.ToFind{
			if lettre_choisi == string(letter){
				h.Word[index] = lettre_choisi
			}else{
				h.Attempts -= 1
			}
		}
	}
}