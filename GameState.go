package main

type GameState struct {
	CommunityCards []Card;
	PlayersInRoom []Player;
	HandsInPlay map[int][]Card;
	BettingRounds []Round;
}


func newGameState() *GameState {
	return &GameState{
		CommunityCards: make([]Card,0,5),
		PlayersInRoom: make([]Player,0,32),
		BettingRounds:   make([]Round,0,3),
		HandsInPlay:      make(map[int][]Card),
	}
}


