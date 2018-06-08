package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)

func main() {

	hub := newHub()
	createdGameState := newGameState()
	dealer := new(Dealer)

	controller := &GameController{createdGameState,hub, dealer}

	hub.gameController = controller
	go hub.run()

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	http.ListenAndServe(":8080",nil)
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


