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

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
