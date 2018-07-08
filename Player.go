package main

type Player struct {
	ChipCount int `json:"chipCount"`
	session UserSession
	PlayerId string `json:"playerId"`
}
