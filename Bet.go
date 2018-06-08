package main

type Bet struct {
	*Action;
	Amount int;
}


func (b Bet) getPlayerId() int {
	return b.Player.Id
}

func (b Bet) getType() string {
	return "bet"
}

func (b Bet) satisfies(requiredBet Bet) bool {
	if b.Amount >= requiredBet.Amount && b.Player.Id == requiredBet.Player.Id {
		return true
	}

	return false
}

