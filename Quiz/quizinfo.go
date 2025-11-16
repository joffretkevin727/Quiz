package main

import "fmt"

func Quizinfo() {

	for i := range Questions_informatique {
		fmt.Println(Questions_informatique[i].Question)
		Answers(Questions_informatique[i].Choice1, Questions_informatique[i].Choice2, Questions_informatique[i].Choice3)
	}

	checkAnswer()
}

func Answers(s string, d string, p string) {
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
		if User_answers[i] == Good_answers_informatique[i] {
			User_score++
		}
	}

	fmt.Println("Vous avez ", User_score, " réponses juste")

	fmt.Println(User_level[User_score])
}
