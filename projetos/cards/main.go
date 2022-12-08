package main

import "fmt"

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("---------------------------------")
	remainingCards.print()

	fmt.Println(cards.toString())

	cards2 := newDeck()
	cards2.saveToFile("my_cards")

	cards3 := readFromFile("my_cards")
	cards3.print()

}

// var card string = "Ace of Spades"
// card2 := "Ace of Spades"   // Initialization
// card2 = "Five of Diamonds" // After initialization

// Type converting
// []byte -> byte slice

// greeting := "Hi there!"
// fmt.Println([]byte(greeting))
// Output: [72 105 32 116 104 101 114 101 33]
// * Values in ASCII code.
