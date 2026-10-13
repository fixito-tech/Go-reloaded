package main

func Apostrophe(mots []string) []string {
	resultat := []string{}
	ouverte := false

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		if mot == "'" {
			if !ouverte && i+1 < len(mots) {

				resultat = append(resultat, "'"+mots[i+1])
				i++
				ouverte = true
				continue
			}
			if ouverte && len(resultat) > 0 {

				dernier := len(resultat) - 1
				resultat[dernier] = resultat[dernier] + "'"
				ouverte = false
				continue
			}
		}

		resultat = append(resultat, mot)
	}

	return resultat
}
