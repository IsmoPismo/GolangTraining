package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardNum := []string{"Ace of ", "Deuce of ", "Three of ", "Four of ", "Five of ", "Six of ", "Seven of ", "Eigth of ", "Nine of", "Ten of ", "Jack of ", "Queen of ", "King of "}
	cardSign := []string{"Hearts", "Diamonts", "Clubs", "Spades"}

	for _, sign := range cardSign {
		for _, num := range cardNum {
			cards = append(cards, num+sign)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, hs int) (deck, deck) {
	return d[:hs], d[hs:]
}
