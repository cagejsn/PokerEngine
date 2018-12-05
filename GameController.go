package main


// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L${SRCDIR} ompeval.a
// #include <stdlib.h>
// #include "handVal.h"
import "C"
import (
	"encoding/json"
	"fmt"
)

type GameController struct {
	masterState   *GameState
	hub           *Hub
	dealer        *Dealer
	stateIsLocked bool
	players       map[string]*Player
}

func (g *GameController) isStateLocked() bool {
	return g.stateIsLocked
}

func (g *GameController) lockState() {
	g.stateIsLocked = true
}

func (g *GameController) unlockState() {
	g.stateIsLocked = false
}

func makePlayerForUserSession(session UserSession) *Player {
	//hash email for the playerId, use ordering of the slice to make sure that the turns go in order
	return &Player{100, session, session.user.Email}
}

func (g *GameController) addUserSession(session UserSession) {
	g.players[session.sessionKey] = makePlayerForUserSession(session)
}

func (g *GameController) removeUserSession(session UserSession) {
	delete(g.players, session.sessionKey)
}

func (g *GameController) processMessage(message []byte) {

		var incomingMessage struct {
			Amount int    `json:"amount"`
			UserSessionKey   string `json:"user"`
			Type   string `json:"type"`
		}

		json.Unmarshal(message, &incomingMessage)

		//get Player for userSession
		playerRef, prs := g.players[incomingMessage.UserSessionKey]
		if !prs {
			fmt.Print("NO PLAYER")
		}


		if incomingMessage.Type == "initialMessage" {
			//deal new game
			if !g.masterState.hasGameStarted() {
				if len(g.players) >= 3 {
					//needs to be passed all the players too
					g.masterState = g.dealer.dealNewGame(&g.players)
				} else {
					return
				}
			}
		}

		g.modifyGameStateFor(Action{incomingMessage.Amount, playerRef, incomingMessage.Type})
		g.hub.outboundState <- *g.masterState

}

func allocateWinnings(bettingRounds []Round, winner *Player){
	for _, round := range bettingRounds {
	
		allocateWinningsForRound(round.Participation,winner)
	}
}

func allocateWinningsForRound(bettingRoundActions Participation, winner *Player){
	for _, action := range bettingRoundActions {

		winner.ChipCount += action.Amount
	}
}

func (g *GameController) modifyGameStateFor(action Action) {

	if g.isStateLocked() {
		return
	} else {
		g.lockState()
		defer g.unlockState()

		gameState := g.masterState
		dealer := g.dealer

		currentBettingRound := gameState.getCurrentBettingRound()

		if !action.satifies(currentBettingRound.RequiredParticipation) {
			return
		}

		previousAction := currentBettingRound.getPreviousActionForPlayer(action.Player.PlayerId)

		action.Player.ChipCount -= (action.Amount - previousAction.Amount)

		previousAction.Type = action.Type
		if previousAction.Type != "fold" {
			previousAction.Amount = action.Amount
		}




		currentBettingRound.RequiredParticipation = nextRequired(&currentBettingRound.Participation)

		//recursive call to modifyGameStateForAction to fold absent player.
		//if _, prs := g.players[currentBettingRound.RequiredParticipation.Player.session.sessionKey]; !prs {
		//	g.modifyGameStateFor(Action{0, currentBettingRound.RequiredParticipation.Player, "fold"})
		//	return
		//}

		for g.masterState.getCurrentBettingRound().isCompleted() {

			switch g.masterState.CurrentRound {
			case 0:
				dealer.dealCommunityCard(gameState)
				dealer.dealCommunityCard(gameState)
				dealer.dealCommunityCard(gameState)
				gameState.BettingRounds = append(gameState.BettingRounds, *newRoundFromParticipation(gameState.BettingRounds[0].Participation))
			case 1:
				dealer.dealCommunityCard(g.masterState)
				gameState.BettingRounds = append(gameState.BettingRounds, *newRoundFromParticipation(gameState.BettingRounds[1].Participation))
			case 2:
				dealer.dealCommunityCard(gameState)
				gameState.BettingRounds = append(gameState.BettingRounds, *newRoundFromParticipation(gameState.BettingRounds[2].Participation))
			case 3:
				//done look for winner
				allocateWinnings(gameState.BettingRounds, findWinner(*g.masterState))
				g.masterState = g.dealer.dealNewGame(&g.players)
				return
			default:
				return
			}
			g.masterState.CurrentRound += 1
		}
	}

}

func findWinner(gameState GameState) *Player {



	communityCards := gameState.CommunityCards
	for _, v := range gameState.BettingRounds[3].Participation {

		handForPlayer := gameState.HandsInPlay[v.Player.PlayerId]



		a := C.malloc(28)
		var b = C.evaluateHand( (*C.int)(a) , 4 )

		print(a, b , communityCards, handForPlayer)
	return v.Player

	}

	return &Player{}
}

type Hand [2]Card

type listOfCard []Card

func (l listOfCard) forEachCombinationInNChooseK(k int, run func([]Card) ){

	var kGroup listOfCard = make(listOfCard,k)

	missingVals := make([]int, len(l) - k )

	var pos int = 0

	for a, _ := range missingVals {
		missingVals[a] = pos
		pos++
	}


	fmt.Print(missingVals)
	fmt.Print("Cage")

	run(kGroup)
}


func findBestHandForPlayer(communityCards [5]Card, hand Hand) [5]Card {

	// var fiveBestCards [5]Card = [5]Card{}
	var sevenCards listOfCard = []Card{
		communityCards[0],
		communityCards[1],
		communityCards[2],
		communityCards[3],
		communityCards[4],
		hand[0],
		hand[1],
	}

	f := func(kList []Card){
		
	


	}

	sevenCards.forEachCombinationInNChooseK(5, f)

  return communityCards
}



