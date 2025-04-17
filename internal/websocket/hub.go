package websocket

import (
	"log"

	"iot-dashboard/internal/device"
)

// Hub manages WebSocket clients
type Hub struct {
	Clients       map[*Client]bool
	Broadcast     chan []byte
	Register      chan *Client
	Unregister    chan *Client
	DeviceManager *device.DeviceManager
	Logger        *log.Logger
}

// NewHub creates a new WebSocket hub
func NewHub(logger *log.Logger, deviceManager *device.DeviceManager) *Hub {
	return &Hub{
		Clients:       make(map[*Client]bool),
		Broadcast:     make(chan []byte),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		DeviceManager: deviceManager,
		Logger:        logger,
	}
}

// Run starts the hub's message handling loop
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.Logger.Printf("Client connected: %s", client.DeviceID)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				h.Logger.Printf("Client disconnected: %s", client.DeviceID)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}