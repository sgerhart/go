package main


func main() {

	//var card string = "Ace of Spades"
	// := only for new vars - initialize
	// card := newCard()
	cards := newDeck()

	hand, remainingCard := deal(cards, 5)

	hand.print()
	remainingCard.print()

	cards.writeToFile("mydeck")


}
