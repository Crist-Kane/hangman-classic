package main

import (
	"fmt"
	"os"
	"time"
)

var mot string = ""
var words string = "test"
var letter string
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

var resultat = []string{}

func (moi *toi) try(text string) {
	var exemple = []rune(words)
	compteur := 0
	compteur2 := 0
	if moi.vieactuel > 0 {
		compteur = 0
		compteur2 = 0
		if text >= "a" && text <= "z" || text >= "A" && text <= "Z" {
			for i := 0; i < len(resultat); i++ {
				for r := 0; r < len(words); r++ {
					if text[0] == words[r] {
						if text == resultat[i] {
							if text == resultat[i] {
								fmt.Println("Dommage c'est la bonne réponse mais elle existe déja")
								moi.choix()
								break
							}
						}
					}
				}
			}
			for i := 0; i < len(resultat); i++ {
				for r := 0; r < len(words); r++ {
					if !(text[0] == words[r]) {
						if text == resultat[i] {
							if text == resultat[i] {
								fmt.Println("Mince alors c'est une mauvaise réponce, mais tu a déja écrit ")
								moi.choix()
								break
							}
						}
					}
				}
			}
			for r := 0; r < len(exemple); r++ {
				if len(text) > 1 && text != words {
					fmt.Println("Attention le programme ne prendra que la première lettre, a moins que tu est trouver le mot au complet")
					time.Sleep(time.Second * 2)
					break
				}
				if text == words {
					fmt.Println("Congratulation tu a gagné")
					time.Sleep(time.Second * 2)
					moi.choix()
					break
				}
				if text[0] == byte(exemple[r]) {
					fmt.Println("Tu as trouver une lettre ")
					moi.choix()
					break
				}
			}
			for r := 0; r < len(exemple); r++ {
				if !(text[0] == byte(exemple[r])) {
					for v := 0; v < len(words); v++ {
						if words[v] != text[0] {
							if compteur2 == 0 {
								fmt.Println("Mince alors ce n'est pas ça !!!")
								time.Sleep(time.Second * 1)
								moi.vieactuel -= 1
								compteur += 1
								fmt.Println("Il te reste ", moi.vieactuel, " PV sur ", moi.viemax)
								time.Sleep(time.Second * 2)
								compteur2 += 1
							}
						}
					}
				}
			}
		}
	}
	resultat = append(resultat, text)
}
func (moi *toi) choix() {
	var test string
	var commande string
	fmt.Println("#------------------------------------#")
	fmt.Println(" Si tu veut arrêter tape fin")
	fmt.Println(" Si tu veut commencer tape choix")
	fmt.Println("#------------------------------------#")
	fmt.Scan(&commande)
	switch commande {
	case "choix":
		fmt.Println("Chose →")
		fmt.Scan(&test)
		moi.try(test)
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
