package main

import (
	"math/rand"
	"time"
)

type Dealer struct {
	deck []Card

}

func (d *Dealer) dealCommunityCard(state *GameState){
	state.CommunityCards = append(state.CommunityCards, d.drawCard())
}

func (d *Dealer) dealNewGame(players *map[string]*Player) *GameState {

	d.deck = *generateNewShuffledDeck()
	state := newGameState()

	for _, player := range *players {

		state.HandsInPlay[player.PlayerId] = make([]Card,2,2)
		state.HandsInPlay[player.PlayerId][0] = d.drawCard()
		state.HandsInPlay[player.PlayerId][1] = d.drawCard()

	}

	state.BettingRounds = append( state.BettingRounds, *newRound(*players))
	return state
}


func (d *Dealer) drawCard() Card {

	length := len(d.deck)

	if length == 0 {
		return Card{}
	}

	returnCard := d.deck[length - 1]
	d.deck = d.deck[0:length - 1]
	return returnCard
}

func generateNewShuffledDeck() *[]Card {

	cards := []Card{

		{2,"H"},
		{2,"C"},
		{2,"S"},
		{2,"D"},

		{3,"H"},
		{3,"C"},
		{3,"S"},
		{3,"D"},

		{4,"H"},
		{4,"C"},
		{4,"S"},
		{4,"D"},

		{5,"H"},
		{5,"C"},
		{5,"S"},
		{5,"D"},

		{6,"H"},
		{6,"C"},
		{6,"S"},
		{6,"D"},

		{7,"H"},
		{7,"C"},
		{7,"S"},
		{7,"D"},

		{8,"H"},
		{8,"C"},
		{8,"S"},
		{8,"D"},

		{9,"H"},
		{9,"C"},
		{9,"S"},
		{9,"D"},

		{10,"H"},
		{10,"C"},
		{10,"S"},
		{10,"D"},

		{11,"H"},
		{11,"C"},
		{11,"S"},
		{11,"D"},

		{12,"H"},
		{12,"C"},
		{12,"S"},
		{12,"D"},

		{13,"H"},
		{13,"C"},
		{13,"S"},
		{13,"D"},

		{1,"H"},
		{1,"C"},
		{1,"S"},
		{1,"D"}}


		deckToFill := make([]Card,52,52)


	for index, _  := range cards {

		rand.Seed(time.Now().UnixNano())
		randomDeckPosition := rand.Int() % len(cards)

		deckToFill[index] = cards[randomDeckPosition]

		//should swap the last card with the just selected one and shorten the slice
		cards[randomDeckPosition] = cards[len(cards) - 1 ]
		cards = cards[0:len(cards) - 1]
	}

	return &deckToFill
}

