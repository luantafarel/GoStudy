package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(path string) error {
	return ioutil.WriteFile(path, []byte(d.toString()), 0666)
}

func newDeckFromFile() deck {
	deckBytes, err := ioutil.ReadFile("./deck.txt")
	if err != nil {
		fmt.Println("Error reading the file " + err.Error())
		fmt.Println("New deck generated")
		/*
			 	fmt.Println("Error reading the file " + err.Error())
				os.Exit(1)
		*/
		return newDeck()
	}
	deckString := strings.Split(string(deckBytes), ",")
	return deck(deckString)
}

func (d deck) shuffleDeck() deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	return d
}
