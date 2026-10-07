package main

import (
	"strconv"
	"strings"
)

func Cap(mots []string) []string {
	resultat := []string{}

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		// cas simple : (cap)
		if mot == "(cap)" {
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = majuscule(resultat[dernier])
			}
			continue
		}

		// cas avec un nombre : "(cap," puis "6)"
		if mot == "(cap," && i+1 < len(mots) {
			nombre, err := strconv.Atoi(strings.TrimSuffix(mots[i+1], ")"))
			if err == nil {
				for j := 1; j <= nombre && len(resultat)-j >= 0; j++ {
					position := len(resultat) - j
					resultat[position] = majuscule(resultat[position])
				}
				i++ // on saute aussi le "6)"
				continue
			}
		}

		resultat = append(resultat, mot)
	}

	return resultat
}

// majuscule met la première lettre en majuscule et le reste en minuscules.
func majuscule(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	premiere := strings.ToUpper(string(lettres[0]))
	return premiere + string(lettres[1:])
}
