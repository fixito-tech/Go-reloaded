package main

//j'ai fais ce programme avec l'IA pour m'aider dans la compréension des différent code.

import "testing"

func TestCorriger(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"hex", "1E (hex) files were added", "30 files were added"},
		{"bin", "It has been 10 (bin) years", "It has been 2 years"},
		{"up", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"low", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"cap", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"up with count", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"count larger than words", "hello world (cap, 10)", "Hello World"},
		{"punctuation", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"ellipsis", "I was thinking ... You were right", "I was thinking... You were right"},
		{"interrobang", "Really !? Yes", "Really!? Yes"},
		{"single quote word", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"quoted sentence", "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"article", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"article h", "a honest man", "an honest man"},
		{"contraction kept", "it's a test", "it's a test"},
		{
			"sample 1",
			"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{"sample 2", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"sample 3", "There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"sample 4", "Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Corriger(tc.in); got != tc.want {
				t.Errorf("\n got: %q\nwant: %q", got, tc.want)
			}
		})
	}
}
