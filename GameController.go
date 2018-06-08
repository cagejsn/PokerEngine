package main

import (
	"log"
	"encoding/json"
	"fmt"
)

type GameController struct {
	masterState *GameState
	hub *Hub
	dealer *Dealer
}


func (g *GameController) addPlayerToRoom(player *Player)  {
	g.masterState.PlayersInRoom = append(g.masterState.PlayersInRoom, *player)
}

func (g *GameController) removePlayerFromRoom(player *Player)  {
}

func (g GameController) processMessage(message []byte){

	action, isAction := processAsAction(message)
	bet, isBet := processAsBet(message)
	fold, isFold := processAsFold(message)

	if !isBet && !isFold && !isAction {
		//return
	}

	if isAction {

		fmt.Print(string(action.Player.Id))

	}



	g.dealer.dealNewGame(g.masterState)
	g.dealer.dealFirstCommunityCards(g.masterState)


	cards := g.masterState.HandsInPlay[1]
	fmt.Print(cards[0])
	fmt.Print(cards[1])



	if isBet {
		currentRound := g.masterState.BettingRounds[0]
		requiredBet := currentRound.nextRequired()

		if bet.satisfies(requiredBet){
			currentRound.Bets[bet.Player.Id] = bet

		}
	}

	if isFold {
		currentRound := g.masterState.BettingRounds[0]
		currentRound.Folds[fold.Player.Id] = fold

	}


	g.hub.newGameState <- *g.masterState


}

func processAsAction(message []byte) (Action, bool) {
	var action Action;
	err := json.Unmarshal(message, &action)

	if err != nil  {
		log.Println(err)
		return action , false
	}



	return action , true
}


func processAsBet(message []byte) (Bet, bool) {
	var bet Bet;

	fmt.Print(string(message))

	err := json.Unmarshal(message, &bet)

	if err != nil  {
		log.Println(err)
		return bet , false
	}
	return bet , true
}


func processAsFold(message []byte) (Fold, bool) {
	var fold Fold;

	err := json.Unmarshal(message, &fold)

	if err != nil || fold.Player.Id == 0 {
		log.Println(err)
		return fold , false
	}
	return fold , true
}

