package quiz

import "fmt"

function ShowMenu() {

    fmt.Println("=====MENU=====")
	fmt.Println("choisissez votre quiz parmis les suivants :")
	fmt.Println(" 1 - Informatique")
	fmt.Println(" 2 - Cybersecurité")
	fmt.Println(" 3 - Ia / Data")

	switch fmt.Scan() {
	case "1":
        quiz.QuizInfo()
	case "2":
        quiz.QuizCyber()
	case "3":
		quiz.QuizIaData()
	}
  
}

