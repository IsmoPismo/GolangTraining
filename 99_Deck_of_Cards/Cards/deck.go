package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())

	// s := time.Now().Nanosecond()
	// fmt.Println(s)
	// fmt.Println(time.Now().Day())
	// fmt.Println(time.Now().Month())
	// fmt.Println(time.Now().Second())
	// e := time.Now().Nanosecond()
	// fmt.Println(e)
	// fmt.Println(e - s)

	r := rand.New(src)

	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}
