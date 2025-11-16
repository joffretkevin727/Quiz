package main

// nombre de réponses juste
var User_score int = 0

// Associe le choix de l'user en int en un string
var Theme = map[int]string{
	1: "Informatique",
	2: "Cybersécurité",
	3: "IA & Data",
}

// enregistrer les choix de l'user ici
var Choice int

//enregistrer les réponses choisies
var User_answers []int

//les bonnes réponses

var Good_answers_info = [10]int{2, 3, 1, 1, 2, 1, 2, 1, 1, 3}
var Good_answers_iadata = [10]int{2, 3, 1, 3, 2, 3, 1, 1, 2, 1}

//Codes ANSI couleurs
var Blue = "\033[34m"
var White = "\033[0m"
var Yellow = "\033[33m"
var Red = "\033[31m"

var User_level = map[int]string{
	1:  "Nobody",
	2:  "Novice",
	3:  "Debutant",
	4:  "Apprenti",
	5:  "Aspirant",
	6:  "Confirmé",
	7:  "Nobody",
	8:  "Aficionados",
	9:  "Pro",
	10: "Expert",
}

type QA struct {
	Question string
	Choice1  string
	Choice2  string
	Choice3  string
}

var Questions_informatique = map[int]QA{
	1:  {"Quel est le type d'une variable qui doit contenir des numeros ?", "var", "int", "string"},
	2:  {"Comment déclare-t-on une fonction en Go ?", "function", "fun", "func"},
	3:  {"Quelle est la sortie de ce code Go ?\nfunc main() {\n    x := 5\n    fmt.Println(&x)\n}", "L'adresse mémoire de x", "La valeur 5", "Une erreur"},
	4:  {"Quel package utilise-t-on pour travailler avec le temps ?", "time", "datetime", "clock"},
	5:  {"Comment déclare-t-on une slice vide en Go ?", "var s []int", "s := []int{}", "var s [0]int"},
	6:  {"Quel mot-clé permet de lancer une goroutine ?", "go", "routine", "goroutine"},
	7:  {"Que fait le mot-clé defer ?", "Exécute la fonction immédiatement", "Exécute la fonction à la fin de la fonction englobante", "Déclare une variable globale"},
	8:  {"Comment lit-on sur un channel ch ?", "<-ch", "ch <- value", "ch.read()"},
	9:  {"Quelle est la différence entre un tableau et une slice ?", "Tableau taille fixe, slice dynamique", "Slice taille fixe, tableau dynamique", "Aucune différence"},
	10: {"Comment crée-t-on une map qui associe des chaînes à des entiers en Go ?", "myMap := map{}", "myMap := map[int]string{}", "myMap := map[string]int{}"},
}
