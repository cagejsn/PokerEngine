package main

type Round struct {
	Participation Participation `json:"participation"`
	RequiredParticipation Action `json:"requiredBet"`
}

type Participation []Action

func newRound(participants map[string]*Player) *Round {

	var roundParticipants = make(Participation,0,32)

	for _, participant := range participants {
		roundParticipants = append(roundParticipants, Action{ 0,participant,"none",})
	}

	return &Round{
		Participation: roundParticipants,
		RequiredParticipation: Action{0,roundParticipants[0].Player,"bet"},
	}
}

//filters out the people who folded on the last round
func newRoundFromParticipation(previousParticipation Participation) *Round {

	var nextRoundParticipation = make(Participation,0,32)
	for _, participation := range previousParticipation {
		if participation.Type != "fold" {
			nextRoundParticipation = append(nextRoundParticipation,Action{ 0,participation.Player,"none"})
		}
	}
	return &Round{
		Participation:nextRoundParticipation,
		RequiredParticipation: nextRequired(&nextRoundParticipation),
	}
}

func (r *Round) getPreviousActionForPlayer(playerId string) *Action {

	for i, action := range r.Participation {
		if playerId == action.Player.PlayerId {
			return &r.Participation[i]
		}
	}

	return nil
}

func (r *Round) isCompleted() bool {

	maxBet := findMaxBet(r.Participation)
	remainingPlayers := len(r.Participation)

	if remainingPlayers == 1 {
		return true
	}

	for _, roundParticipation := range r.Participation {

		switch roundParticipation.Type{

		case "fold":
			remainingPlayers -= 1
		case "bet":
			betAmount := roundParticipation.Amount
			if betAmount < maxBet {
				return false
			}
		case "none":
			return false
		}

		if remainingPlayers == 1 {
			return true
		}
	}

	return true
}


func findMaxBet(bets []Action) int {

	maxSoFar := 0;
	for _ ,v := range bets {
		if v.Amount > maxSoFar {
			maxSoFar = v.Amount
		}
	}
	return maxSoFar;
}

func nextRequired(participation *Participation) Action {

	maxBet := findMaxBet(*participation)

	for _, participation := range *participation {

		switch participation.Type{
		case "fold":
			continue
		case "bet":
			if participation.Amount == maxBet {
				continue
			}
		}

		return Action{maxBet,participation.Player, "bet"}

	}

	return Action{}
}