package main

import (
	"strconv"
	"strings"
)

// Consigne : (low) met le mot d'avant en minuscules.
// (low, 3) met les 3 mots d'avant en minuscules.
// Exemple : "I should stop SHOUTING (low)" -> "I should stop shouting"
func Low(mots []string) []string {
	resultat := []string{}

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		// cas simple : (low)
		if mot == "(low)" {
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = strings.ToLower(resultat[dernier])
			}
			continue
		}

		// cas avec un nombre : "(low," puis "3)"
		if mot == "(low," && i+1 < len(mots) {
			nombre, err := strconv.Atoi(strings.TrimSuffix(mots[i+1], ")"))
			if err == nil {
				for j := 1; j <= nombre && len(resultat)-j >= 0; j++ {
					position := len(resultat) - j
					resultat[position] = strings.ToLower(resultat[position])
				}
				i++ // on saute aussi le "3)"
				continue
			}
		}

		resultat = append(resultat, mot)
	}

	return resultat
}
