package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var a []string = []string{
	"  ##  ",
	" #  # ",
	"######",
	"#    #",
	"#    #",
	"      ",
}
var b []string = []string{
	"##### ",
	"#    #",
	"##### ",
	"#    #",
	"##### ",
	"      ",
}
var c []string = []string{
	" #####",
	"#     ",
	"#     ",
	"#     ",
	" #####",
	"      ",
}
var d []string = []string{
	"##### ",
	"#    #",
	"#    #",
	"#    #",
	"##### ",
	"      ",
}
var e []string = []string{
	"######",
	"#     ",
	"####  ",
	"#     ",
	"######",
	"      ",
}
var f []string = []string{
	"######",
	"#     ",
	"####  ",
	"#     ",
	"#     ",
	"      ",
}
var g []string = []string{
	" #####",
	"#     ",
	"#  ###",
	"#    #",
	" #####",
	"      ",
}
var h []string = []string{
	"#    #",
	"#    #",
	"######",
	"#    #",
	"#    #",
	"      ",
}
var i []string = []string{
	"##### ",
	"  #   ",
	"  #   ",
	"  #   ",
	"##### ",
	"      ",
}
var j []string = []string{
	"     #",
	"     #",
	"     #",
	" #   #",
	"  ##  ",
	"      ",
}
var k []string = []string{
	"#    #",
	"#  #  ",
	"##    ",
	"#  #  ",
	"#    #",
	"      ",
}
var l []string = []string{
	"#     ",
	"#     ",
	"#     ",
	"#     ",
	"######",
	"      ",
}
var m []string = []string{
	"#    #",
	"##  ##",
	"# ## #",
	"#    #",
	"#    #",
	"      ",
}
var n []string = []string{
	"#    #",
	"##   #",
	"# #  #",
	"#  # #",
	"#   ##",
	"      ",
}
var o []string = []string{
	" #### ",
	"#    #",
	"#    #",
	"#    #",
	" #### ",
	"      ",
}
var p []string = []string{
	"##### ",
	"#    #",
	"##### ",
	"#     ",
	"#     ",
	"      ",
}
var q []string = []string{
	" #### ",
	"#    #",
	"#    #",
	" #### ",
	"     # ",
	"      ",
}
var r []string = []string{
	"##### ",
	"#    #",
	"##### ",
	"#  #  ",
	"#    #",
	"      ",
}
var s []string = []string{
	" #####",
	"#     ",
	" #### ",
	"     #",
	"##### ",
	"      ",
}
var t []string = []string{
	"######",
	"  #   ",
	"  #   ",
	"  #   ",
	"  #   ",
	"      ",
}
var u []string = []string{
	"#    #",
	"#    #",
	"#    #",
	"#    #",
	" #### ",
	"      ",
}
var v []string = []string{
	"#    #",
	"#    #",
	" #  # ",
	" #  # ",
	"  ##  ",
	"      ",
}
var w []string = []string{
	"# # # ",
	"# # # ",
	"# # # ",
	"# # # ",
	" # #  ",
	"      ",
}
var x []string = []string{
	"#   # ",
	" # #  ",
	"  #   ",
	" # #  ",
	"#   # ",
	"      ",
}
var y []string = []string{
	"#    #",
	"#    #",
	" #### ",
	"  #   ",
	"  #   ",
	"      ",
}
var z []string = []string{
	"######",
	"    # ",
	"  ##  ",
	" #    ",
	"######",
	"      ",
}
var space []string = []string{
	"      ",
	"      ",
	"      ",
	"      ",
	"      ",
	"      ",
}
var underscore []string = []string{
	"      ",
	"      ",
	"      ",
	"      ",
	"      ",
	"######",
}

var letters [][]string = [][]string{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, space, underscore}

var mot string = ""
var word string = ""
var letter string

type toi struct {
	vieactuel int
	viemax    int
}

func (player *toi) init(vieactuel int, viemax int) {
	player.vieactuel = vieactuel
	player.viemax = viemax
}
func (player *toi) test(w int) {
	if w == 0 {
		fmt.Println("bravo")
	}
}

var guessPlayer string = ""
var Allwordguess rune
var resultat = []string{}
var YouWon bool = false
var count int
var count2 int
var guessplayer string
var guessplayerError = []string{}
var guessPlayerRune rune
var searchwordcount int
var HiddenWord string
var done bool = false

func Show_Hide(word, guessletters string) string {
	final := ""
	for i := 0; i < len(word); i++ {
		if strings.Contains(guessletters, string(word[i])) {
			final += string(word[i])
		} else {
			final += " _ "
		}
	}
	return final
}

var word_give_player string
var ContainAllWord string
var worstLetter string = ""
var GoodLetter string

func (player *toi) HangmanAnnimation() {
	CompteurHangmanRepeat := 0
	if CompteurHangmanRepeat == 0 {
		if player.vieactuel == 10 {
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")

		}
		if player.vieactuel == 9 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")

		}
		if player.vieactuel == 8 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")

		}
		if player.vieactuel == 7 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")
		}
		if player.vieactuel == 6 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf(" /|   |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")
		}
		if player.vieactuel == 5 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf("_/|   |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")
		}
		if player.vieactuel == 4 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf("_/|\\  |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")
		}
		if player.vieactuel == 3 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf("_/|\\_ |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")
		}
		if player.vieactuel == 2 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf("_/|\\_ |\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")
		}
		if player.vieactuel == 1 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf("_/|\\_ |\n")
			fmt.Printf("  |   |\n")
			fmt.Printf(" /    |\n")
			fmt.Printf("      |\n")
			fmt.Printf("      |\n")
			fmt.Printf("========\n")
		}
		if player.vieactuel == 0 {
			fmt.Printf("  +---+\n")
			fmt.Printf("  |   |\n")
			fmt.Printf("  O   |\n")
			fmt.Printf("_/|\\_ |\n")
			fmt.Printf("  |   |\n")
			fmt.Printf(" / \\  |\n")
			fmt.Printf("      |\n")
			fmt.Printf("R.I.P |\n")
			fmt.Printf("========\n")
		}
	}
}

func PrintAscii(text string) {
	upperText := []byte{}
	for i := 0; i < len(text); i++ {
		if text[i] > 96 && text[i] < 123 {
			upperText = append(upperText, text[i]-32-65)
		} else if text[i] == 32 {
			upperText = append(upperText, 26)
		} else if text[i] == 95 {
			upperText = append(upperText, 27)
		} else {
			upperText = append(upperText, text[i]-65)
		}
	}
	asciiIndices := []int{}
	for i := 0; i < len(upperText); i++ {
		asciiIndices = append(asciiIndices, int(upperText[i]))
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < len(asciiIndices); j++ {
			fmt.Print(letters[asciiIndices[j]][i] + " ")
		}
		fmt.Print("\n")
	}

}

var Generate_Random_Word int = 0

func (player *toi) try(word_give_player string) {
	if Generate_Random_Word == 0 {
		rand.Seed(time.Now().UnixNano())
		data, err := ioutil.ReadFile("mot/words.txt")
		if err != nil {
			os.Exit(0)
		}

		words := string(data)
		wordsList := strings.Fields(words)
		word = wordsList[rand.Intn(len(wordsList))]
		Generate_Random_Word = 1
	}
	guessplayer = ""
	if player.vieactuel > 0 {
		var WordGuessRune = []byte(word)
		compteur := 0
		compteur2 := 0
		compteur = 0
		compteur2 = 0
		if count == 0 {
			var Count_Letter_Same int
			count2 = 1
			fmt.Println("Danss le mot générer il y'a ", len(word), " mots a trouvés")
			for i := 0; i < len(resultat); i++ {
				for r := 0; r < len(word); r++ {
					if !(word_give_player[0] == word[r]) {
						if word_give_player == resultat[i] {
							if word_give_player == resultat[i] {
								fmt.Println("Mince alors c'est une mauvaise réponce, mais tu a déja écrit ")
								time.Sleep(time.Second * 1)
								guessplayerError = append(guessplayerError, word_give_player)
								fmt.Println("voici déja utilisé : ", guessplayerError)
								time.Sleep(time.Second * 1)
								fmt.Println(Show_Hide(word, guessPlayer))
								PrintAscii(Show_Hide(word, guessPlayer))
								player.HangmanAnnimation()
								time.Sleep(time.Second * 2)
								player.Game()
								break
							}
						}
					}
				}
			}
			if guessplayer == word {
				fmt.Println("Tu as gagner !!")
				fmt.Println("GG")
				os.Exit(0)
				time.Sleep(time.Second * 2)
			}
			for r := 0; r < len(WordGuessRune); r++ {
				if len(word_give_player) > 1 && word_give_player != word {
					fmt.Println("Ce n'est pas la bonne réponse !!!")
					time.Sleep(time.Second * 1)
					player.vieactuel -= 2
					compteur += 1
					fmt.Println("Il te reste ", player.vieactuel, " PV sur ", player.viemax)
					time.Sleep(time.Second * 1)
					guessplayerError = append(guessplayerError, word_give_player)
					fmt.Println("voici déja utilisé : ", guessplayerError)
					time.Sleep(time.Second * 1)
					fmt.Println(Show_Hide(word, guessPlayer))
					PrintAscii(Show_Hide(word, guessPlayer))
					player.HangmanAnnimation()
					time.Sleep(time.Second * 2)
					break
				}
				if word_give_player == word {
					fmt.Println("Congratulation tu a gagné")
					fmt.Println(Show_Hide(word, guessPlayer))
					PrintAscii(Show_Hide(word, guessPlayer))
					time.Sleep(time.Second * 2)
					Generate_Random_Word = 0
					player.Game()
					break
				}
			}
			if word_give_player == word {
				fmt.Println("Tu as gagner")
				fmt.Println(Show_Hide(word, guessPlayer))
				Generate_Random_Word = 0
				player.Game()
			}
			if len(word_give_player) == 1 {
				if guessplayer == word {
					fmt.Println("Tu as gagner !!")
					fmt.Println(Show_Hide(word, guessPlayer))
					fmt.Println("GG")
					Generate_Random_Word = 0
					player.Game()
					time.Sleep(time.Second * 2)
				}
				for r := 0; r < len(WordGuessRune); r++ {
					if word_give_player[0] == byte(WordGuessRune[r]) {
						if word_give_player == guessPlayer {
							fmt.Println("C'est une bonne réponse hélas tu ne peut pas car tu la déja donner")
							player.Game()
							break
						} else {
							fmt.Println("Tu as trouver une lettre ")
							Count_Letter_Same++
							guessPlayer += word_give_player
							time.Sleep(time.Second * 1)
							guessplayerError = append(guessplayerError, word_give_player)
							fmt.Println("voici déja utilisé : ", guessplayerError)
							time.Sleep(time.Second * 1)
							fmt.Println(Show_Hide(word, guessPlayer))
							PrintAscii(Show_Hide(word, guessPlayer))
							player.HangmanAnnimation()
							if Show_Hide(word, guessPlayer) == Show_Hide(word, "abcdefghijklmnopqrstuvwxyz") {
								PrintAscii(Show_Hide(word, guessPlayer))
								fmt.Println("Vous avez Gagné!")
								fmt.Printf("Le Mot Etait %s\n", word)
								time.Sleep(time.Second * 1)
								Generate_Random_Word = 0
								player.Game()
							}
							time.Sleep(time.Second * 2)
							fmt.Println(string(Allwordguess))
							player.Game()
							break
						}
					}
				}
				for r := 0; r < len(WordGuessRune); r++ {
					if !(word_give_player[0] == byte(WordGuessRune[r])) {
						for v := 0; v < len(word); v++ {
							if word[v] != word_give_player[0] {
								if compteur2 == 0 {
									fmt.Println("Mince alors ce n'est pas ça !!!")
									time.Sleep(time.Second * 1)
									player.vieactuel -= 1
									compteur += 1
									fmt.Println("Il te reste ", player.vieactuel, " PV sur ", player.viemax)
									time.Sleep(time.Second * 2)
									compteur2 += 1
									worstLetter += word_give_player
									time.Sleep(time.Second * 1)
									guessplayerError = append(guessplayerError, word_give_player)
									fmt.Println("voici déja utilisé : ", guessplayerError)
									fmt.Println(Show_Hide(word, guessPlayer))
									PrintAscii(Show_Hide(word, guessPlayer))
									player.HangmanAnnimation()
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
			fmt.Println("La réponse était : ", word)
			player.vieactuel = 10
			player.Game()
		}
		if Show_Hide(word, guessPlayer) == Show_Hide(word, "abcdefghijklmnopqrstuvwxyz") {
			fmt.Println("Vous avez Gagné!")
			fmt.Printf("Le Mot Etait %s\n", word)
			time.Sleep(time.Second * 1)
			Generate_Random_Word = 0
			player.Game()
		}
	}
	resultat = append(resultat, word_give_player)
}
func Win(guessPlayer string, word string) bool {
	if guessPlayer == word {
		return true
	}
	return false
}
func (player *toi) NoLetter(word_give_player string) {
	if len(worstLetter) >= 1 {
		fmt.Println("Les mauvaises lettres sont : ", worstLetter)
		time.Sleep(time.Second * 2)
		player.Game()
	} else {
		fmt.Println("Aucune mauvaises lettre INCROYABLE, tu a vraiment jouer au moins ???")
		time.Sleep(time.Second * 2)
		player.Game()
	}
}
func (player *toi) GoodLetterFindYou(word_give_player string) {
	if strings.Contains(word, word_give_player) {
		if len(guessPlayer) >= 1 {
			fmt.Println("Les bonnes lettres sont : ", guessPlayer)
			time.Sleep(time.Second * 2)
			player.Game()
		} else {
			fmt.Println("Tu n'a pas encore trouver de bonne lettre ")
			time.Sleep(time.Second * 2)
			player.Game()
		}
	}
}
func (player *toi) Game() {
	var test string
	var commande string
	fmt.Println("#----------------------------------------------------------------------------#")
	fmt.Println("| Si tu veut arrêter tape fin                                                |")
	fmt.Println("| Si tu veut commencer tape choix                                            |")
	fmt.Println("| Si tu veut savoir quel mauvaise lettre tu a déja trouver : Tapez → 2       |")
	fmt.Println("| Si tu veut savoir quel bonne lettre tu a déja trouver : Tapez → 1          |")
	fmt.Println("#----------------------------------------------------------------------------#")
	fmt.Scan(&commande)
	switch commande {
	case "choix":
		fmt.Println("Chose →")
		fmt.Scan(&test)
		player.try(test)
		player.Game()
	case "2":
		player.NoLetter(test)
		player.Game()
		//renvoie les bonnes lettres et les mauvaisses lettres dissocier
	case "1":
		player.GoodLetterFindYou(test)
		player.Game()
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