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
var r = []rune("test")

type toi struct {
	vieactuel int
	viemax    int
}

func (moi *toi) init(vieactuel int, viemax int) {
	moi.vieactuel = vieactuel
	moi.viemax = viemax

}
func (moi *toi) choix() {
	var commande string
	fmt.Println("Si tu veut recommencer tape fin")
	fmt.Println("Chose →")
	fmt.Scan(&commande)
	switch commande {
	case string(letter):
		fmt.Scan(&letter)
		for i := 0; i < len(r); i++ {
			if moi.vieactuel > 0 {
				if letter == r[i] {
					fmt.Println("Bravo tu a trouver une lettre")
					time.Sleep(2)
					moi.choix()
				} else {
					if letter != r[i] {
						fmt.Println("Tu na pas trouver la bonne lettre")
						moi.vieactuel -= 1
						if moi.vieactuel <= 0 {
							fmt.Println("Tu as perdue")
							os.Exit(0)
						}
						moi.choix()
					}
				}
			}
		}

	case "fin":
		os.Exit(0)
	}
}
func random() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(words); i++ {
	}
}
func main() {
	fmt.Println("ceci est un test")
	random()

}
