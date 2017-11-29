package controller

import (
	"socket"
)

func RegisterRoutes(h *socket.Hub) {
	hub = h
	new(socketController).RegisterRoute()
}
