package main

type Fold struct {
	*Action
}

func (f Fold) getPlayerId() int {
	return f.Player.Id;
}


func (f Fold) getType() string {
	return "fold"
}

