// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type StateController interface {
	addUserSession(session UserSession)
	removeUserSession(session UserSession)
	processMessage([]byte)
}

type State interface {
	customizeStateForClient(session UserSession) State
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	clientActions chan []byte

	// Registered clients.
	clients map[*Client]bool

	// outbound gameState to the Clients
	outboundState chan State

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	stateController StateController;
}

func newHub() *Hub {
	return &Hub{
		clientActions: make(chan []byte),
		outboundState: make(chan State),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:

			h.clients[client] = true
			h.stateController.addUserSession(client.userSession)


		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {

				delete(h.clients, client)
				close(client.send)

				h.stateController.removeUserSession(client.userSession)
			}

		case message := <- h.clientActions:
			go h.stateController.processMessage(message)

		case message := <-h.outboundState:
			for client := range h.clients {

				sendableGameState := message.customizeStateForClient(client.userSession)

				select {
				case client.send <- sendableGameState:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

