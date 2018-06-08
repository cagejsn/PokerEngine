package main

import (
	"math/rand"
	"time"
)

type Dealer struct {
	deck []Card

}

func (d *Dealer) dealFirstCommunityCards(state *GameState){

	state.CommunityCards = append(state.CommunityCards, d.drawCard())
	state.CommunityCards = append(state.CommunityCards, d.drawCard())
	state.CommunityCards = append(state.CommunityCards, d.drawCard())

}

func (d *Dealer) dealNewGame(state *GameState) {

	d.deck = *generateNewShuffledDeck()

	for _, player := range state.PlayersInRoom {


		//playerCards := make([]*Card,2,2)
		//playerCardsOne := d.drawCard()
		//playerCardsTwo := d.drawCard()
		state.HandsInPlay[player.Id] = make([]Card,2,2)
		state.HandsInPlay[player.Id][0] = d.drawCard()
		state.HandsInPlay[player.Id][1] = d.drawCard()


		//state.HandsInPlay[player.Id] = []Card{ d.drawCard(), d.drawCard()};
	}

	state.BettingRounds = append(state.BettingRounds,*newRound(state.PlayersInRoom))
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

		{2,"Hearts"},
		{2,"Clubs"},
		{2,"Spades"},
		{2,"Diamonds"},

		{3,"Hearts"},
		{3,"Clubs"},
		{3,"Spades"},
		{3,"Diamonds"},

		{4,"Hearts"},
		{4,"Clubs"},
		{4,"Spades"},
		{4,"Diamonds"},

		{5,"Hearts"},
		{5,"Clubs"},
		{5,"Spades"},
		{5,"Diamonds"},

		{6,"Hearts"},
		{6,"Clubs"},
		{6,"Spades"},
		{6,"Diamonds"},

		{7,"Hearts"},
		{7,"Clubs"},
		{7,"Spades"},
		{7,"Diamonds"},

		{8,"Hearts"},
		{8,"Clubs"},
		{8,"Spades"},
		{8,"Diamonds"},

		{9,"Hearts"},
		{9,"Clubs"},
		{9,"Spades"},
		{9,"Diamonds"},

		{10,"Hearts"},
		{10,"Clubs"},
		{10,"Spades"},
		{10,"Diamonds"},

		{11,"Hearts"},
		{11,"Clubs"},
		{11,"Spades"},
		{11,"Diamonds"},

		{12,"Hearts"},
		{12,"Clubs"},
		{12,"Spades"},
		{12,"Diamonds"},

		{13,"Hearts"},
		{13,"Clubs"},
		{13,"Spades"},
		{13,"Diamonds"},

		{1,"Hearts"},
		{1,"Clubs"},
		{1,"Spades"},
		{1,"Diamonds"}}


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

