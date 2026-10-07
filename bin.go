package main

import "strconv"


func Bin(mots []string) []string {
	resultat := []string{}

	for _, mot := range mots {
		if mot == "(bin)" {
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				nombre, err := strconv.ParseInt(resultat[dernier], 2, 64)
				if err == nil {
					resultat[dernier] = strconv.FormatInt(nombre, 10)
				}
			}
			continue // on ne garde pas le "(bin)"
		}
		resultat = append(resultat, mot)
	}

	return resultat
}
