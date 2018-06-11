package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/kabukky/httpscerts"
	"log"
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

	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8080")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)

}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


