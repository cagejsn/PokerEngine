package main

type Round struct {
	Participants []Player
	currentPlayerTurn int
	Bets map[int]Bet
	Folds map[int]Fold
}

func newRound(participants []Player) *Round {
	return &Round{
		Participants: participants,
		currentPlayerTurn: 0,
		Bets: make(map[int]Bet),
		Folds: make(map[int]Fold),

	}
}

func (r Round) isCompleted() bool {



	var highestBet int = findMaxBet(r.Bets);

	var equalsHighestBet = func(bet Bet) bool {
		if bet.Amount == highestBet {
			return true;
		}
		return false;
	}

	for _ , player := range r.Participants {

		_ , hasFolded := r.Folds[player.Id]
		if hasFolded {
			continue
		} else {

			bet, hasBet := r.Bets[player.Id]
			if !hasBet {
				return false;
			}

			if(!equalsHighestBet(bet)){
				return false;
			}
		}
	}

	return true;
}


func findMaxBet(bets map[int]Bet) int {

	maxSoFar := 0;
	for _ ,v := range bets {
		if v.Amount > maxSoFar {
			maxSoFar = v.Amount
		}
	}
	return maxSoFar;
}

func (r *Round) nextRequired() Bet {

	for _, player := range r.Participants {

		bet, hasBet := r.Bets[player.Id]
		_, hasFolded := r.Folds[player.Id]

		if hasFolded {
			continue
		}

		maxBet := findMaxBet(r.Bets)
		if hasBet {
			if bet.Amount == maxBet {
				continue
			}
		}

		return Bet{
			&Action{player, "bet"},maxBet}

	}
	return Bet{}
}