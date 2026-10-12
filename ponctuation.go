package main

import "strings"

func Ponctuation(mots []string) []string {
	resultat := []string{}

	for _, mot := range mots {

		nbSignes := 0
		for nbSignes < len(mot) && strings.ContainsRune(".,!?:;", rune(mot[nbSignes])) {
			nbSignes++
		}

		if nbSignes > 0 && len(resultat) > 0 {

			dernier := len(resultat) - 1
			resultat[dernier] = resultat[dernier] + mot[:nbSignes]

			reste := mot[nbSignes:]
			if reste != "" {
				resultat = append(resultat, reste)
			}
			continue
		}

		resultat = append(resultat, mot)
	}

	return resultat
}
