package main

// import "fmt"

var deckSize int

func main() {
	// var card string
	// card = "Ace of Spades"
	// card := "Ace of Spades"
	// fmt.Println(card)

	// card := newCard()
	// fmt.Println(card)

	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
	// fmt.Println("\n\n...Shuffling...\n\n")
	cards.shuffle()
	cards.print()
}
