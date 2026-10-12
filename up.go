package main

import (
	"strconv"
	"strings"
)

func Up(mots []string) []string {
	resultat := []string{}

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		if mot == "(up)" {
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = strings.ToUpper(resultat[dernier])
			}
			continue
		}

		if mot == "(up," && i+1 < len(mots) {
			nombre, err := strconv.Atoi(strings.TrimSuffix(mots[i+1], ")"))
			if err == nil {
				for j := 1; j <= nombre && len(resultat)-j >= 0; j++ {
					position := len(resultat) - j
					resultat[position] = strings.ToUpper(resultat[position])
				}
				i++
				continue
			}
		}

		resultat = append(resultat, mot)
	}

	return resultat
}
