package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

//-------------Générer un nombre aléatoire----------------\\

func Nombres_aléatoire(min, max int) int {
	return min + rand.Intn(max-min)
}

//-------------Transformer le mot rand en mot à deviner----------------\\

func deviner(s string) string {
	var mot []string      // EX : salut
	var motcaché []string // EX : _ _ _ _ _
	for _, u := range s {
		mot = append(mot, string(u))
		motcaché = append(motcaché, "_")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < (len(mot))/2-1; i++ {
		pos := rand.Intn(len(mot))
		motcaché[pos] = mot[pos]
	}
	vide := ""
	for i := 0; i < len(motcaché); i++ {
		vide = vide + motcaché[i] + " " // pour pas avoir de OUT OF RANGE
	}
	return vide
}

//-------Remplacer '_' par la lettre saisie-------\\

func remplace_lettre(mot string, caché string, lettresrentrée string) string {
	var tg []rune
	ouai := []rune(lettresrentrée)
	vide := ""
	for _, s := range mot {
		tg = append(tg, s)
	}
	for i, s2 := range caché {
		if s2 == '_' && tg[i] == ouai[0] {
			s2 = tg[i]
		}
		vide = vide + string(s2)
	}
	return vide
}

//--------------LETTRE SAISIE---------------\\

func Lettre_Trouvé(lettre string, word string) bool { //savoir si la lettre saisie est bien dans le mot qui est caché
	for _, motentier := range word { // on parcours le string
		if string(motentier) == lettre { // Si dans le string il y a la lettre ou un mot on return true sinon false
			return true
		}
	}
	return false
}

//--------------MOT SAISIE---------------\\

func mot_trouvé(mot_rentrée string, mot string) bool {
	if mot_rentrée == mot {
		return true
	}
	return false
}

//--------------Rajouter des espaces dans le mot----------------\\

func add_space(bj string) string {
	slice := ""
	for a, b := range bj {
		if a != len(bj) {
			slice = slice + string(b) + " "
		} else {
			slice = slice + string(b)
		}
	}
	return slice
}

//--------------Toutes les lettres sont-elles dans le mot ?----------------\\

func toutes_lettres(mot string, caché string) bool {
	for _, tg := range caché {
		tqt := 0
		for _, tg2 := range mot {
			if string(tg) == string(tg2) {
				tqt = 1
			}
		}
		if tqt == 0 {
			return false
		}
	}
	return true
}

//--------------Changer les lettres MIN en lettres MAJ (isupper)----------------\\

func maj(mot string) string {
	tqt := []rune(mot)
	for i := range tqt {
		if tqt[i] >= 'a' && tqt[i] <= 'z' {
			tqt[i] = tqt[i] - 32
		}
	}
	return string(tqt)
}

//--------------STOCKER LES LETTRES UTILISEE----------------\\

func lettre_stockée(Int_user string, tab []string) bool {
	if len(tab) == 0 { //si il n'y a pas de lettre
		return true
	} else { //sinon s'il y a une lettre
		for _, w := range tab { // on parcours la lettre
			if Int_user == string(w) { // si cette lettre est = à la lettre rentrée (si elle a déjà été uitilisé)
				return false
			}
		}
	}
	return true
}

//--------------position du pendu en fonction des attempts----------------\\

func Pendu(r int) string {
	if r == 9 {
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("=========\n")
	}
	if r == 8 {
		fmt.Printf("       \n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if r == 7 {
		fmt.Printf("  +---+\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if r == 6 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if r == 5 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if r == 4 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if r == 3 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if r == 2 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\ |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if r == 1 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf(" /    |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if r == 0 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf(" / \\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	return string(r)
}

//----------Structure du jeux------------\\

type José struct {
	attempts int
}

func main() {

	//----------iINITIALISATION DES VARIABLES DE LA STRUCT------------\\

	var g1 José
	g1.attempts = 10

	var slice []string //création d'un slice
	var tab []string   //création d'un slice pour stockée les lettres déjà utilisés

	//----------Ouvrir le fichier .txt------------\\

	file, error := os.Open("words.txt") //ouverture du fichier
	if error != nil {                   //définir l'err si il n'y a pas de fichié
		log.Fatal(error)
	}
	fileScanner := bufio.NewScanner(file) //Scanner le fichier
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() { //Mettre chaque ligne du fichier dans un slice vide
		slice = append(slice, (fileScanner.Text()))
	}

	//----------Initialisation des variables------------\\

	max := len(slice) //valeurs max et min du slice
	min := 0
	rand.Seed(time.Now().UnixNano())
	rand := (Nombres_aléatoire(min, max)) //chiffre random par rapport à la longeur du fichier
	mot := slice[rand]                    //Mot aléatoire
	caché := deviner(mot)                 //défini le mot caché
	fmt.Println("GOOD LUCK, YOU HAVE 10 ATTEMPTS")
	nar := maj(caché) //mot chaché en majuscule
	fmt.Println(nar)
	qz := add_space(mot)
	for g1.attempts != 0 || caché == qz {

		//----------Scanner la lettre saisie par l'utilisateur------------\\

		scan := bufio.NewScanner(os.Stdin) // Permet de scan la saisie de l'utilisateur
		fmt.Println("CHOOSE:  ")
		scan.Scan()             //On scan la saisie
		Int_user := scan.Text() //On prend la lettre saisi pour la mettre dans une variable   //On prend le mot saisi pour la mettre dans une variable

		//----------L'UTILISATEUR SAISIE UN MOT------------\\
		if len(Int_user) != 0 {
			if len(Int_user) > 1 {
				abc := mot_trouvé(mot, Int_user)
				if abc == true { //SI LE MOT SAISIE EST LE BON
					fmt.Println("CONGRATS !")
					fmt.Println("THE WORD WAS", maj(mot))
					break
				} else { // SI CE N'EST PAS LE BON
					if g1.attempts <= 1 {
						fmt.Println("YOU LOOSE")
						fmt.Println("THE WORD WAS", maj(mot))
						break
					} else {
						g1.attempts = g1.attempts - 2
						fmt.Println("THIS IS THE WRONG WORD")
					}
				}

				//----------STOCKER LES LETTRES UTILISEES------------\\

			} else {
				if len(Int_user) == 1 { //si l'intput est une seule lettre
					ui := lettre_stockée(Int_user, tab)
					if ui == true {
						tab = append(tab, Int_user)
					} else {
						fmt.Println("ERROR, LETTER WAS ALREADY USED")
					}
				}
			}
		}
		if Lettre_Trouvé(Int_user, mot) == true {
			ok := remplace_lettre(qz, caché, Int_user)
			bienb := maj(ok)
			caché = bienb
			fmt.Println(bienb)
			z := maj(qz)
			e := maj(caché)
			if toutes_lettres(z, e) == true { // Si toutes les lettres ont été trouvés
				fmt.Println("CONGRATS !")
				break // casse la fonction main (arrête tout)
			} else {
				continue
			}
		} else {
			if Lettre_Trouvé(Int_user, mot) == false && len(Int_user) < 2 {
				g1.attempts-- //Si il ya des err
			}
			if g1.attempts > 0 {
				fmt.Println("THE LETTER IS NOT IN THE WORD :", g1.attempts, "REMAINING") // S'il reste encore des essais
			} else {
				fmt.Println("YOU LOOSE")
				MOT := maj(mot)
				fmt.Println("THE WORD WAS :", MOT)
			}
			g := Pendu(g1.attempts) //Définir la position du pendu en fonction du nombre d'essaies
			a := remplace_lettre(qz, caché, Int_user)
			l := maj(a)
			fmt.Println(l)
			fmt.Println(g)
		}
	}
}
