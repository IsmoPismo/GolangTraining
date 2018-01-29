package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	hand, _ := deal(cards, 2)
	fmt.Println(hand)
}
