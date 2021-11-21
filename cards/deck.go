package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type call 'deck'
// which is a slice of strings
type deck []string

// testing two
func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"Diamonds", "Hearts", "Spades", "clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) print() {

	for i, card := range d {

		fmt.Println(i, card)

	}

}

// Converts string slice to string
func (d deck) toString() string {

	newString := strings.Join(d, ",")

	return newString

}

func (d deck) writeToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}
