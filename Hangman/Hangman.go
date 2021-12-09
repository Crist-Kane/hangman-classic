package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)
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
var underscore string = " _ "
var YouWon bool = false
var count int
var count2 int
var guessplayer string
var guessplayerError = []string{}
var guessPlayerRune rune
var searchwordcount int
var HiddenWord string
var done bool = false
var final string = ""

func Show_Hide(word, word_give_player string) string {
	final = ""
	for i := 0; i < len(word); i++ {
		if strings.Contains(word_give_player, string(word[i])) {
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

func (player *toi) You_Win() {
	fmt.Println("     ###############         ###       ###    ############     ###       ####")
	fmt.Println("     ###############         ###       ###    ############     ###       ####")
	fmt.Println("     ###############         ###       ###    ############     ###       ####")
	fmt.Println("       ###########           ###       ###    ###      ###     ###       ####")
	fmt.Println("       ###########           ###       ###    ###      ###     ###       ####")
	fmt.Println("       ###########           ###       ###    ###      ###     ###       ####")
	fmt.Println("    ##################       ###      ####    ###      ###     ###       ####")
	fmt.Println("   ####################      ####     ####    ###      ###     ###       ####")
	fmt.Println("   ####################      #############    ###      ###     ###       ####")
	fmt.Println("   ###  ########  #####      #############    ###      ###     ###       ####")
	fmt.Println("   ###  ########  #####      #############    ###      ###     ###       ####")
	fmt.Println("   ###   #######  #####        #########      ###      ###     ###       ####")
	fmt.Println("  ###    #######     ###          ###         ###      ###     ###       ####")
	fmt.Println("  ###    #######     ###          ###         ###      ###     ###       ####")
	fmt.Println(" ###      ######      ###         ###         ###      ###     ###       ####")
	fmt.Println(" ###       #####      ###         ###         ###      ###     ###       ####")
	fmt.Println("  ###       ####     ###          ###         ###      ###     ###       ####")
	fmt.Println("  ###       ####     ###          ###         ###      ###     ###       ####")
	fmt.Println("   ###   #######  #####           ###         ###      ###     ###       ####")
	fmt.Println("   ###   #######  #####           ###         ###      ###     ###       ####")
	fmt.Println("   ###################            ###         ###      ###     ###       ####")
	fmt.Println("    #### ###### #####             ###         ###      ###     ###       ####")
	fmt.Println("    #### ###### #####             ###         ###      ###     ###       ####")
	fmt.Println("         ######                   ###         ###      ###     ###       ####")
	fmt.Println("         ######                   ###         ###      ###     ###       ####")
	fmt.Println("         ######                   ###         ###      ###     ###       ####")
	fmt.Println("         ######                   ###         ###      ###     ###       ####")
	fmt.Println("         ######                   ###         ###      ###     ###       ####")
	fmt.Println("        ########                  ###         ###      ###     ###       ####")
	fmt.Println("        ########                  ###         ############     ##############")
	fmt.Println("      ############                ###         ############     ##############")
	fmt.Println("    ################              ###         ############      ############ ")
	fmt.Println("    ################              ###         ############      ############ ")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("                                                                       #####")
	fmt.Println("                                                                       #####")
	fmt.Println(" #####             ####   ###################     #####                #####")
	fmt.Println(" #####             ####   ###################     #####                #####")
	fmt.Println(" #####             ####   ###################     ######               #####")
	fmt.Println(" #####             ####   ###################     #######              #####")
	fmt.Println(" #####             ####         #######           #######              #####")
	fmt.Println(" #####             ####         #######           ########             #####")
	fmt.Println(" #####             ####         #######           ########             #####")
	fmt.Println(" #####             ####         #######           #########            #####")
	fmt.Println(" #####     ###     ####         #######           ###########          #####")
	fmt.Println(" #####     ###     ####         #######           ###########          #####")
	fmt.Println(" #####     ###     ####         #######           #### #######         #####")
	fmt.Println(" #####    #####    ####         #######           #### #######         #####")
	fmt.Println(" #####    #####    ####         #######           ####   ######        #####")
	fmt.Println(" #####   #######   ####         #######           ####    #######      #####")
	fmt.Println(" #####   #######   ####         #######           ####    #######      #####")
	fmt.Println(" #####  ####   ### ####         #######           ####      ################")
	fmt.Println(" #####  ####   ### ####         #######           ####      ################")
	fmt.Println(" #########      #######         #######           ####        ##############")
	fmt.Println(" #########      #######         #######           ####         #############")
	fmt.Println(" #######         ######         #######           ####           ###########")
	fmt.Println(" #######         ######         #######           ####           ###########")
	fmt.Println(" ######           #####         #######           ####             #########")
	fmt.Println(" ######           #####         #######           ####              ########")
	fmt.Println(" #####             ####   ###################     ####                 #####")
	fmt.Println(" ####               ###   ###################     ####                 #####")
	fmt.Println(" ####               ###   ###################     ####                 #####")
	fmt.Println(" ####               ###   ###################     ####                 #####")
	time.Sleep(time.Second * 2)
}

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
			if !(word_give_player == word) {
				//fmt.Println("Le ot chacher est : ", word) //si tu ne trouve pas le mot l'avoir directemnt
				fmt.Println("Danss le mot générer il y'a ", len(word), " mots a trouvés")
			}
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
				final = ""
				worstLetter = ""
				guessPlayer = ""
				player.You_Win()
				player.vieactuel = 10
				Generate_Random_Word = 0
				guessplayerError = []string{}
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
					player.HangmanAnnimation()
					time.Sleep(time.Second * 2)
					break
				}
				if word_give_player == word {
					fmt.Println("Congratulation tu a gagné")
					time.Sleep(time.Second * 2)
					final = ""
					worstLetter = ""
					guessPlayer = ""
					player.You_Win()
					player.vieactuel = 10
					Generate_Random_Word = 0
					guessplayerError = []string{}
					player.Game()
					break
				}
			}
			if word_give_player == word {
				fmt.Println("Tu as gagner")
				final = ""
				worstLetter = ""
				guessPlayer = ""
				player.You_Win()
				player.vieactuel = 10
				guessplayerError = []string{}
				Generate_Random_Word = 0
				player.Game()
			}
			if len(word_give_player) == 1 {
				if guessplayer == word {
					fmt.Println("Tu as gagner !!")
					fmt.Println("GG")
					final = ""
					worstLetter = ""
					guessPlayer = ""
					player.You_Win()
					player.vieactuel = 10
					guessplayerError = []string{}
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
							player.HangmanAnnimation()
							if Show_Hide(word, guessPlayer) == Show_Hide(word, "abcdefghijklmnopqrstuvwxyz") {
								fmt.Println("Vous avez Gagné!")
								fmt.Printf("Le Mot Etait %s\n", word)
								final = ""
								worstLetter = ""
								guessPlayer = ""
								player.You_Win()
								player.vieactuel = 10
								guessplayerError = []string{}
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
			final = ""
			worstLetter = ""
			guessPlayer = ""
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("                      :::!~!!!!!:.")
			fmt.Println("                  .xUHWH!! !!?M88WHX:.")
			fmt.Println("                .X*#M@$!!  !X!M$$$$$$WWx:.")
			fmt.Println("               :!!!!!!?H! :!$!$$$$$$$$$$8X:")
			fmt.Println("              !!~  ~:~!! :~!$!#$$$$$$$$$$8X:")
			fmt.Println("             ~!~!!!!~~ .:XW$$$U!!?$$$$$$RMM!")
			fmt.Println("               !:~~~ .:!M'T#$$$$WX??#MRRMMM!")
			fmt.Println("               ~?WuxiW*`   `'#$$$$8!!!!??!!!")
			fmt.Println("             :X- M$$$$       `'T#$T~!8$WUXU~")
			fmt.Println("            :%`  ~#$$$m:        ~!~ ?$$$$$$")
			fmt.Println("          :!`.-   ~T$$$$8xx.  .xWW- ~''##*'")
			fmt.Println(".....   -~~:<` !    ~?T#$$@@W@*?$$      /`")
			fmt.Println("W$@@M!!! .!~~ !!     .:XUW$W!~ `'~:    :")
			fmt.Println("#'~~`.:x%`!!  !H:   !WM$$$$Ti.: .!WUn+!`")
			fmt.Println(":::~:!!`:X~ .: ?H.!u '$$$B$$$!W:U!T$$M~")
			fmt.Println(".~~   :X@!.-~   ?@WTWo('*$$$W$TH$! `")
			fmt.Println("Wi.~!X$?!-~    : ?$$$B$Wu('**$RM!")
			fmt.Println("$R@i.~~ !     :   ~$$$$$B$$en:``")
			fmt.Println("?MXT@Wx.~    :     ~'##*$$$$M~")
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			time.Sleep(time.Second * 2)
			guessplayerError = []string{}
			Generate_Random_Word = 0
			player.vieactuel = 10
			player.Game()
		}
		if Show_Hide(word, guessPlayer) == Show_Hide(word, "abcdefghijklmnopqrstuvwxyz") {
			fmt.Println("Vous avez Gagné!")
			fmt.Printf("Le Mot Etait %s\n", word)
			final = ""
			worstLetter = ""
			guessPlayer = ""
			player.You_Win()
			player.vieactuel = 10
			guessplayerError = []string{}
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
