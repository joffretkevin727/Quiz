package main

import "fmt"

func Quizinfo() {

	fmt.Println("Quel est le type d'une variable qui doit contenir des numeros")
	answers("var", "int", "string")

	fmt.Println("Comment declare t-on une fonction en go ?")
	answers("function", "fun", "func")

	fmt.Println("Quelle est la sortie de code Go ?")
	fmt.Print(Blue, "func", Yellow, " main() {\n", White, "	x := 5\n	fmt.Println(&x)\n}\n")
	answers("L'addresse mémoire de x", "La valeur 5", "une erreur")

	fmt.Println("Quel package utilise t-on pour travailler avec le temps ?")
	answers("time", "datetime", "clock")

	fmt.Println("Comment déclare-t-on une slice vide en Go ?")
	answers("var s []int", "s := []int{}", "var s [0]int")

	fmt.Println("Quel mot-clé permet de lancer une goroutine ?")
	answers("go", "routine", "goroutine")

	fmt.Println("Que fait le mot-clé defer ?")
	answers("Exécute la fonction immédiatement", "Exécute la fonction à la fin de la fonction englobante", "Déclare une variable globale")

	fmt.Println("Comment lit-on sur un channel ch ?")
	answers("<-ch", "ch <- value", "ch.read()")

	fmt.Println("Quelle est la différence entre un tableau et une slice ?")
	answers("Tableau taille fixe, slice dynamique", "Slice taille fixe, tableau dynamique", "Aucune différence")

	fmt.Println("Comment crée-t-on une map qui associe des chaînes à des entiers en Go ?")
	answers("myMap := map{}", "myMap := map[int]string{}", "myMap := map[string]int{}")

	checkAnswer()
}

func answers(s string, d string, p string) {
	fmt.Print("1. ", s, "\n2. ", d, "\n3. ", p, "\n")
	limitChoice()
}

func limitChoice() {
	fmt.Scan(&Choice)

	for Choice > 3 || Choice < 1 {
		fmt.Println("Entrez un chiffre compris entre 1 et 3")
		fmt.Scan(&Choice)
	}
	User_answers = append(User_answers, Choice)
}

func checkAnswer() {
	for i := range User_answers {
		if User_answers[i] == Good_answers[i] {
			User_score++
		}
	}

	fmt.Println("Vous avez ", User_score, " réponses juste")

	fmt.Println(User_level[User_score])
}
