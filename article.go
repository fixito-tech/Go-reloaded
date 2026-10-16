package main

import "strings"

func Article(mots []string) []string {
	for i := 0; i < len(mots)-1; i++ {
		if mots[i] != "a" && mots[i] != "A" {
			continue
		}

		suivant := strings.TrimLeft(mots[i+1], "'")
		if suivant == "" {
			continue
		}

		premiereLettre := strings.ToLower(string(suivant[0]))
		if strings.Contains("aeiouh", premiereLettre) {
			mots[i] = mots[i] + "n"
		}
	}

	return mots
}
