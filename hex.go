package main

import "strconv"

func Hex(mots []string) []string {
	resultat := []string{}

	for _, mot := range mots {
		if mot == "(hex)" {
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				nombre, err := strconv.ParseInt(resultat[dernier], 16, 64)
				if err == nil {
					resultat[dernier] = strconv.FormatInt(nombre, 10)
				}
			}
			continue
		}
		resultat = append(resultat, mot)
	}

	return resultat
}
