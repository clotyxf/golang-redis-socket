package socket

import (
	"log"
)

type Hub struct {
	clients map[*Client]bool

	//broadcast chan []byte
	broadcast chan *PMessage

	subscribe chan *Client

	unsubscribe chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:     make(map[*Client]bool),
		broadcast:   make(chan *PMessage),
		subscribe:   make(chan *Client),
		unsubscribe: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.subscribe:
			h.clients[client] = true
		case client := <-h.unsubscribe:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				log.Println(client)
				select {
				case client.send <- message.data:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
