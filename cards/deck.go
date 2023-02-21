package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

// Create a new type of 'deck' 
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string {"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + "-" + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() string{
	return "Shuffled"
}

// Takes deck and handSize as arguments
// Returns two values; both of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Returns complete string representation of the given deck
func (d deck) deckToString() string {
	// res := ""
	// for _, value := range []string(d) {
	// 	res = res + "," + value
	// }

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte (d.deckToString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}