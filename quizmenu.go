package quiz

import "fmt"

func Quizmenu() {
	fmt.Println("Bienvenue sur le quiz")

	fmt.Println("Bienvenue sur le quizz informatique")
	fmt.Println("Choisissez le theme")
	fmt.Print("1. Informatique\n2. Cybersecurite\n3. IAData\n")

	for !IsQuizzOver {

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
			Quiziadata()
		case 3:
			Quizcyber()
		}

	}
}
