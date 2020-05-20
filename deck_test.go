package main

import (
	"os"
	"testing"
)

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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52 , but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	amountOfCards := 5
	d := newDeck()

	hand, remainingDeck := d.deal(amountOfCards)

	if len(hand) != amountOfCards {
		t.Errorf("Expected %d cards in hand but got %v", amountOfCards, len(hand))
	}
	if len(remainingDeck) != len(d)-len(hand) {
		t.Errorf("Expected %d cards in the remaining deck but got %v", amountOfCards, len(remainingDeck))
	}

	for id, remainingCard := range remainingDeck {
		for ih, handCard := range hand {
			if handCard == remainingCard {
				t.Errorf("Expected no duplicated cards in hand and remaning deck, but found %v in hand position %d and remaining deck position %d", handCard, id, ih)
			}
		}
	}
}
