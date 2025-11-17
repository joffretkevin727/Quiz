package main

func Quizcyber() {
	User_answers = []int{}
	User_score = 0

	// Intro "Matrix / Terminal"
	fmt.Println(Red + "\n🛑  SYSTEM ALERT : PROTOCOLE DE SÉCURITÉ AVANCÉ  🛑" + White)
	fmt.Println(Red + "==========================================================" + White)
	fmt.Println(Yellow + "💻  ANALYSE DES COMPÉTENCES TECHNIQUES...  💻" + White)
	fmt.Println(Blue + "➡️   Niveau : Intermédiaire (Pas de droit à l'erreur)" + White)
	fmt.Println(Red + "==========================================================\n" + White)

	// Q1 - Typosquatting
	fmt.Println(Yellow + "🔗  1. Vous voulez aller sur Facebook. Lequel de ces liens est une tentative de 'Typosquatting' (URL piège) ?" + White)
	answers(
		"www.facebook.com",
		"www.faceboook.com (3 'o')",
		"m.facebook.com",
	)

	// Q2 - Hashing vs Encryption
	fmt.Println(Yellow + "\n🔐  2. Comment un site web sécurisé doit-il stocker votre mot de passe dans sa base de données ?" + White)
	answers(
		"En clair (Texte brut) pour pouvoir vous le renvoyer si vous l'oubliez.",
		"Chiffré avec une clé symétrique (réversible).",
		"Haché (Hash) avec une fonction irréversible (ex: Argon2 ou SHA-256).",
	)

	// Q3 - MITM
	fmt.Println(Yellow + "\n📡  3. Sur un Wi-Fi public non sécurisé, une attaque 'Man-in-the-Middle' permet à un pirate de :" + White)
	answers(
		"Intercepter et modifier les données échangées entre vous et le site web.",
		"Prendre le contrôle total de votre processeur.",
		"Supprimer votre compte Facebook à distance.",
	)

	// Q4 - Ransomware (Technique)
	fmt.Println(Yellow + "\n💸  4. Quelle est la particularité technique d'un Ransomware comparé à un virus classique ?" + White)
	answers(
		"Il ralentit le système en minant des cryptomonnaies.",
		"Il utilise le chiffrement (cryptographie) pour rendre vos fichiers inaccessibles sans la clé.",
		"Il se replique uniquement par email.",
	)

	// Q5 - Zero-Day
	fmt.Println(Yellow + "\n🐞  5. Qu'est-ce qu'une faille 'Zero-Day' ?" + White)
	answers(
		"Une faille qui n'existe que depuis 0 jour (aujourd'hui).",
		"Une faille qui ne touche que les logiciels gratuits.",
		"Une vulnérabilité non connue de l'éditeur (pas encore de correctif disponible) exploitée par des pirates.",
	)

	// Q6 - VPN (Réalité)
	fmt.Println(Yellow + "\n🌍  6. Que cache réellement un VPN aux yeux de votre fournisseur d'accès Internet (FAI) ?" + White)
	answers(
		"Le contenu de votre trafic et la destination finale (le site visité).",
		"Le fait que vous utilisez un ordinateur ou un téléphone.",
		"Absolument tout, même votre consommation de données.",
	)

	// Q7 - Phishing (Analyse)
	fmt.Println(Yellow + "\n📧  7. Dans un email professionnel, quel indice doit immédiatement vous alerter d'une attaque (Spear Phishing) ?" + White)
	answers(
		"Le mail contient des fautes d'orthographe.",
		"L'adresse de l'expéditeur est 'support@micros0ft.com' (Notez le zéro).",
		"Le mail a été envoyé à 18h00.",
	)

	// Q8 - DDoS
	fmt.Println(Yellow + "\n🔥  8. Une attaque DDoS (Déni de service distribué) consiste à :" + White)
	answers(
		"Voler les mots de passe de tous les utilisateurs.",
		"Effacer la base de données du serveur.",
		"Submerger un serveur avec des milliers de fausses requêtes pour le rendre inaccessible.",
	)

	// Q9 - Social Engineering
	fmt.Println(Yellow + "\n🧠  9. Le principe de 'l'urgence' en ingénierie sociale sert à :" + White)
	answers(
		"Court-circuiter la réflexion critique de la victime pour la forcer à agir vite.",
		"Tester la rapidité du réseau informatique.",
		"Vérifier si la victime est bien devant son écran.",
	)

	// Q10 - Cookies
	fmt.Println(Yellow + "\n🍪  10. Si un pirate vole votre 'Cookie de Session', que peut-il faire ?" + White)
	answers(
		"Connaître votre historique de recherche Google.",
		"Accéder à votre compte connecté sans avoir besoin de votre mot de passe.",
		"Rien, les cookies ne servent qu'à la publicité.",
	)

	checkAnswerCyber()
}

func checkAnswerCyber() {
	for i := range User_answers {
		if User_answers[i] == Good_answers_cyber[i] {
			User_score++
		}
	}

	fmt.Println(Red + "\n================================================" + White)
	fmt.Println("       💾  RAPPORT D'AUDIT  💾       ")
	fmt.Print("       🏆 SCORE : ", User_score, " / 10\n")

	if User_score <= 4 {
		fmt.Println(Red + "\n❌  RÉSULTAT : VULNÉRABLE  ❌" + White)
		fmt.Println("Vos connaissances techniques sont insuffisantes face aux menaces modernes.")
	} else if User_score <= 7 {
		fmt.Println(Yellow + "\n⚠️   RÉSULTAT : AVERTIS  ⚠️" + White)
		fmt.Println("Bonnes bases, mais attention aux détails techniques pointus.")
	} else {
		fmt.Println(Blue + "\n🛡️   RÉSULTAT : OPERATEUR SÉCURITÉ  🛡️" + White)
		fmt.Println("Excellent. Vous comprenez les mécanismes d'attaque et de défense.")
	}

	fmt.Println(Red + "\n   Niveau technique : " + User_level[User_score] + White)
	fmt.Println(Red + "================================================" + White)
}

