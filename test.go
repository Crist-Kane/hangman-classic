package main

import (
	"fmt"
	"os"
	"time"
)

var mot string = ""
var words string = "test"
var letter string
var r = []rune(words)
var exmple = []rune{}

type toi struct {
	vieactuel int
	viemax    int
}

func (moi *toi) init(vieactuel int, viemax int) {
	moi.vieactuel = vieactuel
	moi.viemax = viemax
}
func (moi *toi) test(w int) {
	if w == 0 {
		fmt.Println("bravo")
	}
}
func (moi *toi) help(t rune) {
	fmt.Println("marche")
	if t == 'a' {
		fmt.Println("j'eassaye")
		for e := 0; e < len(r); e++ {
			if t == exmple[e] {
				fmt.Println("Tu as déja donner cette lettre")
			} else {
				if !(t == exmple[e]) {
					for i := 0; i < len(r); i++ {
						if moi.vieactuel > 0 {
							if t == r[i] {
								fmt.Println("Bravo tu a trouver une lettre")
								time.Sleep(2)
								moi.choix()
							} else {
								if t != r[i] {
									fmt.Println("Tu na pas trouver la bonne lettre")
									moi.vieactuel -= 1
									fmt.Println("Il te reste ", moi.vieactuel, " Pv")
									time.Sleep(2)
									if moi.vieactuel <= 0 {
										fmt.Println("Tu as perdue")
										os.Exit(0)
									}
									moi.choix()
								}
							}
						}
					}

				}
			}
		}
	}
}
func (moi *toi) choix() {
	var test rune
	var commande string
	fmt.Println("Si tu veut arrêter tape fin")
	fmt.Println("Si tu veut commencer tape choix")
	fmt.Scan(&commande)
	switch commande {
	case "choix":
		fmt.Println("Chose →")
		fmt.Scan(&test)
		moi.help(test)
		moi.choix()
	case "fin":
		os.Exit(0)
	}
}

var test toi

func main() {
	fmt.Println("ceci est un test")
	test.init(10, 10)
	test.choix()

}
