package main

import (
	"os"
	"strings"
	"testing"
)

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func TestNewDeck(t *testing.T) {
	var typesArr []string
	card := newDeck()
	if len(card) != 52 {
		t.Errorf("Missing cards, expected 50 but got %d", len(card))
	}
	for _, cardType := range []string{"Spades", "Clubs", "Hearts", "Diamonds"} {
		typesArr = filter(card, func(s string) bool {
			return strings.HasSuffix(s, cardType)
		})
		if len(typesArr) != (52 / 4) {
			t.Errorf("Missing cards, expected %d but got %d", 52/4, len(typesArr))
		}
	}
}

func TestSaveFile(t *testing.T) {
	path := "_deckTesting.txt"
	os.Remove(path)
	card := newDeck()
	card.saveToFile(path)
	_, err := os.Stat(path)
	if err != nil {
		t.Errorf("File was not created correctly")
	}
	os.Remove(path)
}
