package main

// import "fmt"

func main()  {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
}