package main

var deckSize int

func main() {
	// var card string
	// card = "Ace of Spades"
	// card := "Ace of Spades"
	// fmt.Println(card)

	// card := newCard()
	// fmt.Println(card)

	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
