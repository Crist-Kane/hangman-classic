package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var mot string = ""
var words string = "test"
var letter string

type toi struct {
	vieactuel int
	viemax    int
}

const Barre_Retour_HangMan = string(byte(92))

func (player *toi) init(vieactuel int, viemax int) {
	player.vieactuel = vieactuel
	player.viemax = viemax
}
func (player *toi) test(w int) {
	if w == 0 {
		fmt.Println("bravo")
	}
}

var Allwordguess rune
var AnnimationHangman = []string{}
var resultat = []string{}
var hidden_result_guess string
var underscore string = " _ "
var count int
var count2 int
var guessplayer string
var guessplayerError = []string{}
var guessPlayerRune rune
var searchwordcount int

func (player *toi) try(word_give_player string) {
	//si byte de words vaut byte de resultat vaut true si chaque lettre corespond au mot que je chercher
	var WordGuessRune = []byte(words)
	CompteurHangmanRepeat := 0
	compteur := 0
	compteur2 := 0
	if player.vieactuel > 0 {
		if CompteurHangmanRepeat == 0 {
			if player.vieactuel == 10 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n         \n      |  \n      |  \n      |  \n      |  \n      |  \n=========")
			}
			if player.vieactuel == 9 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n         \n      |  \n      |  \n      |  \n      |  \n      |  \n=========")
			}
			if player.vieactuel == 8 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n         \n      |  \n      |  \n      |  \n      |  \n      |  \n=========")
			}
			if player.vieactuel == 7 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========")
			}
			if player.vieactuel == 6 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========")
			}
			if player.vieactuel == 5 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========")
			}
			if player.vieactuel == 4 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n  |   |  \n  O   |  \n /|"+Barre_Retour_HangMan+"  |  \n      |  \n      |  \n=========")
			}
			if player.vieactuel == 3 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n  |   |  \n  O   |  \n /|"+Barre_Retour_HangMan+"  |  \n /    |  \n      |  \n=========")
			}
			if player.vieactuel == 2 {
				AnnimationHangman = []string{}
				AnnimationHangman = append(AnnimationHangman, "\n  |   |  \n  O   |  \n /|"+Barre_Retour_HangMan+"  |  \n / "+Barre_Retour_HangMan+"  |  \n      |  \n=========")
			}
		}
		CompteurHangmanRepeat = 1
		compteur = 0
		compteur2 = 0
		if count == 0 {
			var Count_Letter_Same int
			for n := 0; n < len(words); n++ {
				if count2 == 0 {
					hidden_result_guess += underscore
				}
			}
			count2 = 1
			for i := 0; i < len(resultat); i++ {
				for r := 0; r < len(words); r++ {
					if !(word_give_player[0] == words[r]) {
						if word_give_player == resultat[i] {
							if word_give_player == resultat[i] {
								fmt.Println("Mince alors c'est une mauvaise réponce, mais tu a déja écrit ")
								time.Sleep(time.Second * 1)
								guessplayerError = append(guessplayerError, word_give_player)
								fmt.Println("voici déja utilisé : ", guessplayerError)
								time.Sleep(time.Second * 1)
								fmt.Println(AnnimationHangman)
								time.Sleep(time.Second * 2)
								player.Game()
								break
							}
						}
					}
				}
			}

			if guessplayer == words {
				fmt.Println("Tu as gagner !!")
				fmt.Println("GG")
				player.Game()
				time.Sleep(time.Second * 2)
			}
			for r := 0; r < len(WordGuessRune); r++ {
				if len(word_give_player) > 1 && word_give_player != words {
					fmt.Println(hidden_result_guess)
					fmt.Println("Ce n'est pas la bonne réponse !!!")
					time.Sleep(time.Second * 1)
					player.vieactuel -= 2
					compteur += 1
					fmt.Println("Il te reste ", player.vieactuel, " PV sur ", player.viemax)
					time.Sleep(time.Second * 1)
					guessplayerError = append(guessplayerError, word_give_player)
					fmt.Println("voici déja utilisé : ", guessplayerError)
					time.Sleep(time.Second * 1)
					fmt.Println(AnnimationHangman)

					time.Sleep(time.Second * 2)
					break
				}
				if word_give_player == words {
					fmt.Println("Congratulation tu a gagné")
					time.Sleep(time.Second * 2)
					player.Game()
					break
				}
			}
			if string(hidden_result_guess) == words {
				fmt.Println("Tu as gagner")
			}
			if len(word_give_player) == 1 {
				if guessplayer == words {
					fmt.Println("Tu as gagner !!")
					fmt.Println("GG")
					player.Game()
					time.Sleep(time.Second * 2)
				}
				for r := 0; r < len(WordGuessRune); r++ {
					if word_give_player[0] == byte(WordGuessRune[r]) {
						fmt.Println("Tu as trouver une lettre ")
						if strings.Contains(words, word_give_player) {
							if word_give_player[0] == words[0] {
								hidden_result_guess = strings.Replace(hidden_result_guess, "_", word_give_player, 1)
								searchwordcount++

							}
							for j := 0; j < len(words); j++ {
								if word_give_player[0] == words[1] {
									if searchwordcount == 1 {
										hidden_result_guess = strings.Replace(hidden_result_guess, "_", word_give_player, 1)
										searchwordcount++
									}
								}
								if len(words) == j+1 {
									if word_give_player[0] == words[2] {
										if searchwordcount == 2 {
											hidden_result_guess = strings.Replace(hidden_result_guess, "_", word_give_player, 1)
										}
									}
								}
								if len(words) == j+1 {
									if word_give_player[0] == words[3] {
										if searchwordcount == 3 {
											hidden_result_guess = strings.Replace(hidden_result_guess, "_", word_give_player, 1)
											searchwordcount++
										}
									}
								}
								if len(words) == j+1 {
									if word_give_player[0] == words[j] {
										if searchwordcount == 4 {
											hidden_result_guess = strings.Replace(hidden_result_guess, "_", word_give_player, 1)
											searchwordcount++
										}
									}
								}
							}
						}

						Count_Letter_Same++
						time.Sleep(time.Second * 1)
						guessplayerError = append(guessplayerError, word_give_player)
						fmt.Println("voici déja utilisé : ", guessplayerError)
						fmt.Println(hidden_result_guess)
						time.Sleep(time.Second * 1)
						fmt.Println(AnnimationHangman)
						time.Sleep(time.Second * 2)
						fmt.Println(string(Allwordguess))
						player.Game()
						break
					}
				}
				for r := 0; r < len(WordGuessRune); r++ {
					if !(word_give_player[0] == byte(WordGuessRune[r])) {
						for v := 0; v < len(words); v++ {
							if words[v] != word_give_player[0] {
								if compteur2 == 0 {
									fmt.Println(hidden_result_guess)
									fmt.Println("Mince alors ce n'est pas ça !!!")
									time.Sleep(time.Second * 1)
									player.vieactuel -= 1
									compteur += 1
									fmt.Println("Il te reste ", player.vieactuel, " PV sur ", player.viemax)
									time.Sleep(time.Second * 2)
									compteur2 += 1
									time.Sleep(time.Second * 1)
									guessplayerError = append(guessplayerError, word_give_player)
									fmt.Println("voici déja utilisé : ", guessplayerError)
									time.Sleep(time.Second * 1)
									fmt.Println(AnnimationHangman)
									time.Sleep(time.Second * 2)
								}
							}
						}
					}
				}
			}
		}
		if player.vieactuel <= 0 {
			fmt.Println("Tu as PERDUE!!!")
			fmt.Println("La réponse était : ", words)
			player.vieactuel = 10
			player.Game()
		}
	}
	resultat = append(resultat, word_give_player)
}
func (player *toi) Game() {
	var test string
	var commande string
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println("| Si tu veut arrêter tape fin                                        |")
	fmt.Println("| Si tu veut commencer tape choix                                    |")
	fmt.Println("| Si tu veut savoir quel lettre tu a déja utiliser : Tapez → 0       |")
	fmt.Println("| Donner ma réponse : Tapez → 1                                      |")
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Scan(&commande)
	switch commande {
	case "choix":
		fmt.Println("Chose →")
		fmt.Scan(&test)
		player.try(test)
		player.Game()
	case "0":
		//renvoie les bonnes lettres et les mauvaisses lettres dissocier
	case "1":
	case "fin":
		os.Exit(0)
	}
}

var test toi

func main() {
	fmt.Println("ceci est un test")
	test.init(10, 10)
	test.Game()

}
