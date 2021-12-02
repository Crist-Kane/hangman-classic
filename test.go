package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var mot string = ""
var words = []string{"banc", "exmple"}
var letter rune

func choix() {
	var commande string
	fmt.Println("Si tu veut recommencer tape fin")
	fmt.Println("Chose →")
	fmt.Scan(&commande)
	switch commande {
	case letter:
		if letter == words {
			fmt.Scan(&letter)
			charCreation(letter)
			choix()
		}

	case "fin":
		os.Exit(0)
	}
}
func random() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(words); i++ {
		random := rand.Intn(len(words))
	}
}
func main() {
	fmt.Println("ceci est un test")
	choix()
	random()

}
