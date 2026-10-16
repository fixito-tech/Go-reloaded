package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Utilisation : go run . sample.txt result.txt")
		return
	}

	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	lignes := strings.Split(string(contenu), "\n")
	for i := 0; i < len(lignes); i++ {
		lignes[i] = Corriger(lignes[i])
	}
	texte := strings.Join(lignes, "\n")

	err = os.WriteFile(os.Args[2], []byte(texte), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}
}

func Corriger(ligne string) string {
	mots := strings.Fields(ligne)

	mots = Hex(mots)
	mots = Bin(mots)
	mots = Up(mots)
	mots = Low(mots)
	mots = Cap(mots)
	mots = Ponctuation(mots)
	mots = Apostrophe(mots)
	mots = Article(mots)

	return strings.Join(mots, " ")
}
