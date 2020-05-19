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

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
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

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Something went wrong!")
		fmt.Println(err)
		os.Exit(-1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		randomPosition := r.Intn(len(d) - 1)
		d[i], d[randomPosition] = d[randomPosition], d[i]
	}
}
