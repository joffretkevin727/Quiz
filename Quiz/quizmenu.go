package main

import "fmt"

func ShowMenu() {
	fmt.Println("\n====================================================")
	fmt.Println("                Bienvenue sur le quiz !               ")
	fmt.Println("             Choisissez le theme du quiz :            ")
	fmt.Print("     1. Informatique\n2. Cybersécurité\n3. IA & Data\n  ")
	fmt.Println("\n====================================================")

	for !IsQuizOver {

		fmt.Scan(&Choice)

		for Choice > 3 || Choice < 1 {
			fmt.Println("Entrez un chiffre entre 1 et 3 pour choisir le theme")
			fmt.Scan(&Theme)
		}

		fmt.Println("Vous avez choisi le theme", Theme[Choice])

		switch Choice {
		case 1:
			Quizinfo()
		case 2:
			Quizcyber()
		case 3:
			Quiziadata()
		}

	}
}
