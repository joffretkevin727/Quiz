package main

import "fmt"

func Quizinfo() {

	fmt.Println("1. Quel est le type d'une variable qui doit contenir des numeros")
	answers("var", "int", "string")
	fmt.Println("2. Comment declare t-on une fonction en go ?")
	answers("function", "fun", "func")
	fmt.Println("3. Comment declare t-on une fonction en go ?")
	answers("function", "fun", "func")
	fmt.Println("4. Comment declare t-on une fonction en go ?")
	answers("function", "fun", "func")
	checkAnswer()
	fmt.Println("Vous avez ", User_score, " réponses juste")
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
}
