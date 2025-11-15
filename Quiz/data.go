package main

// nombre de réponses juste
var User_score int = 0

// Associe le choix de l'user en int en un string
var Theme = map[int]string{
	1: "Informatique",
	2: "Cybersecurite",
	3: "IA",
}

// enregistrer les choix de l'user ici
var Choice int

//enregistrer les réponses choisies
var User_answers []int

//les bonnes réponses
var Good_answers = [10]int{2, 3, 1, 1, 2, 1, 2, 1, 1, 3}

//Codes ANSI couleurs
var Blue = "\033[34m"
var White = "\033[0m"
var Yellow = "\033[33m"
var Red = "\033[31m"
