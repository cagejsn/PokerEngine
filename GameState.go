package main

type GameState struct {
	CommunityCards []Card `json:"communityCards"`
	HandsInPlay    map[string][]Card `json:"handsInPlay"'`
	BettingRounds  []Round `json:"bettingRounds"`
	CurrentRound   int `json:"currentRound"`
}


func newGameState() *GameState {
	return &GameState{
		CommunityCards: make([]Card, 0, 5),
		BettingRounds:  make([]Round, 0, 3),
		HandsInPlay:    make(map[string][]Card),
	}
}

func (g *GameState) getCurrentBettingRound() *Round {
	return &g.BettingRounds[g.CurrentRound]
}


func (g *GameState) hasGameStarted() bool {

	if len(g.BettingRounds) == 0 {
		return false
	}

	return true
}

//todo move this to GameController
func (g GameState) customizeStateForClient(session UserSession) State {

	gameStateForClient := g
	gameStateForClient.HandsInPlay = make(map[string][]Card)
	for k, v := range g.HandsInPlay {
		if k == session.user.Email {
			gameStateForClient.HandsInPlay[k] = v
		} else {
			gameStateForClient.HandsInPlay[k] = []Card{{0, "X"}, {0, "X"}}
		}
	}

	return gameStateForClient
}
