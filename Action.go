package main

type Action struct {
	Amount int `json:"amount"`
	Player *Player `json:"player"`
	Type string `json:"type"`
}

func (a *Action) satifies(requiredAction Action) bool {
	if a.Player.PlayerId != requiredAction.Player.PlayerId {
		return false
	}

	if a.Type == "fold" {
		return true
	}

	if a.Type ==  "bet" {
		if a.Amount >= requiredAction.Amount {
			return true
		}
		return false
	}

	return false
}
