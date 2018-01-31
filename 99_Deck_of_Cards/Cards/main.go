package main

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 2)
	hand.saveToFile("Ace and Deuce")
	file := newDeckFromFile("Ace and Deuce")
	file.print()
}
