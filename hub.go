// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	clientActions chan []byte

	// Registered clients.
	clients map[*Client]bool

	// outbound gameState to the Clients
	newGameState chan GameState

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	gameController *GameController;
}

func newHub() *Hub {
	return &Hub{
		clientActions: make(chan []byte),
		newGameState: make(chan GameState),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.gameController.addPlayerToRoom(client.Player)


		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {

				delete(h.clients, client)
				close(client.send)

				//h.gameController.removePlayerFromRoom(client)
			}

		case message := <- h.clientActions:
			go h.gameController.processMessage(message)

		case message := <-h.newGameState:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}