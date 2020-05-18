package main

import "fmt"

func main() {
	cards := newDeck()
	hand, cards2 := cards.deal(3)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining deck:")
	cards2.print()
}
