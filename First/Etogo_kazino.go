package main

import "fmt"
import "math/rand"

type Card struct {
	Value rune
	Suit string
}

func CreateCards() [52]Card {
	suit := []string{
		"Diamonds",
		"Hearts",
		"Spades",
		"Clubs",
	}

	var CardsFromKiosk = [52]Card{}
	currentCard := 0
	for i := 0; i<4; i++ {
		for q := 2; q<10; q++ {
			CardsFromKiosk[currentCard].Value = '0' + rune(q)
			CardsFromKiosk[currentCard].Suit = suit[i]
			currentCard++;
		}
		CardsFromKiosk[currentCard].Value = '0'
		CardsFromKiosk[currentCard].Suit = suit[i]
		currentCard++

		CardsFromKiosk[currentCard].Value = 'J'
		CardsFromKiosk[currentCard].Suit = suit[i]
		currentCard++

		CardsFromKiosk[currentCard].Value = 'Q'
		CardsFromKiosk[currentCard].Suit = suit[i]
		currentCard++

		CardsFromKiosk[currentCard].Value = 'K'
		CardsFromKiosk[currentCard].Suit = suit[i]
		currentCard++

		CardsFromKiosk[currentCard].Value = 'A'
		CardsFromKiosk[currentCard].Suit = suit[i]
		currentCard++

	}
	return CardsFromKiosk;

}

func PoryadokKard(card *[52]Card)  {
	swapcard := Card{}
	for i := 0;  i<52; i++ {
		randomI := rand.Intn(51 - 0 + 1) + 0
		swapcard =  (*card)[i]
		(*card)[i] = (*card)[randomI]
		(*card)[randomI] = swapcard
	}

}




func main()  {

	cards := CreateCards()
	for i := 0; i < 52; i++ {
		fmt.Printf("Card %c with suit %s \n", cards[i].Value, cards[i].Suit)
	}

	fmt.Println("\nAfter shuffle : ")
	PoryadokKard(&cards)
	for i := 0; i < 52; i++ {
		fmt.Printf("Card %c with suit %s \n", cards[i].Value, cards[i].Suit)
	}





}
