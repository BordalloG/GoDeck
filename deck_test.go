package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 , but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}

	for i, card := range d {
		for l, card2 := range d {
			if card == card2 && i != l {
				t.Errorf("Expected no duplicated cards, but got %v in position %d and %d", d[i], i, l)
			}
		}
	}
}
