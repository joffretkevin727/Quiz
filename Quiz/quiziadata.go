package main

import "fmt"

func Quiziadata() {
	fmt.Println("Quel est l'objectif principal de l'Intelligence Artificielle (IA) en tant que domaine d'étude ?")
	answers("Créer des robots physiques capables de remplacer l'homme.", "Développer des systèmes informatiques capables de simuler l'intelligence humaine pour effectuer des tâches.", "Programmer des ordinateurs pour effectuer uniquement des tâches répétitives et simples.")

	fmt.Println("Quel terme désigne le sous-domaine de l'IA qui permet aux systèmes d'apprendre automatiquement à partir de données sans être explicitement programmés ?")
	answers("Le Cloud Computing", "La Cybersécurité", "Le Machine Learning")

	fmt.Println("Comment appelle-t-on une IA conçue et entraînée pour effectuer une tâche spécifique unique, comme jouer aux échecs ou prédire la météo ?")
	answers("L'IA Étroite (ou Faible)", "L'IA Générale (ou Forte)", "L'IA Méta")

	fmt.Println("Dans le cadre du Machine Learning, quelle catégorie d'apprentissage utilise des données étiquetées (c'est-à-dire avec les réponses déjà fournies) pour entraîner le modèle ?")
	answers("L'Apprentissage Non Supervisé", "L'Apprentissage par Renforcement", "L'Apprentissage Supervisé")

	fmt.Println("Quel test historique, proposé par Alan Turing, a pour but de déterminer si une machine peut présenter un comportement intelligent indiscernable de celui d'un humain ?")
	answers("Le Test de Bêta", "Le Test de Turing", "Le Test de Kolmogorov")

	fmt.Println("Le Deep Learning (Apprentissage Profond) est un type de Machine Learning qui utilise des architectures composées de plusieurs couches. Comment nomme-t-on ces architectures ?")
	answers("Les Microprocesseurs", "Les Bases de Données Relationnelles", "Les Réseaux Neuronaux")

	fmt.Println("Quelle application de l'IA permet aux machines d'interpréter et de comprendre la parole ou le texte humain, comme Siri ou Google Traduction ?")
	answers("Le Traitement Automatique du Langage Naturel (TALN)", "La Vision par Ordinateur (Computer Vision)", "L'Analyse de Données")

	fmt.Println("Lors d'un Apprentissage Non Supervisé, quel est l'objectif principal du modèle, étant donné que les données ne sont pas étiquetées ?")
	answers("Trouver des structures, des motifs ou des regroupements (clustering) dans les données.", "Prédire une valeur numérique continue (régression)", "Classifier de nouvelles données dans des catégories existantes.")

	fmt.Println("Qu'est-ce que l'apprentissage par renforcement (Reinforcement Learning) ?")
	answers("Un type d'apprentissage où le modèle est formé par des paires d'entrées et de sorties déjà validées.", "Un type d'apprentissage où un agent interagit avec un environnement pour apprendre la meilleure stratégie par essais et erreurs.", "Un type d'apprentissage qui se concentre uniquement sur la découverte des anomalies dans les données.")

	fmt.Println("Quel est le rôle crucial des données d'entraînement ('training data') pour un modèle de Machine Learning ?")
	answers("Elles fournissent l'expérience à partir de laquelle le modèle ajuste ses paramètres internes pour devenir plus précis.", "Elles servent uniquement à valider la vitesse d'exécution du modèle, et non sa précision.", "Elles sont principalement utilisées pour stocker le modèle après son entraînement.")

}
