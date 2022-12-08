package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	// First return  : index '0' up to index handSize (not included, exclusive)
	// Second return : index handSize (inclusive) up to last index (inclusive)
	// Syntax example: d[0:2], d[:2], d[2:]

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	cards := []string(d)
	return strings.Join(cards, ",") // Converts slice into a single string with values separated by ','.
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {

	bs, err := os.ReadFile(filename)

	if err != nil {
		// Option #1 - Log the error and return a call to newDeck().
		// Option #2 - Log the error and entirely quit the program.
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Converts bs (byte slice) -> string (single) -> split string -> slice of strings -> deck type
	return deck([]string(strings.Split(string(bs), ",")))
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano()) // Using time for seed source
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // Random function between 0 and len(d) - 1
		d[i], d[newPosition] = d[newPosition], d[i] // Swap the elements
	}
}
