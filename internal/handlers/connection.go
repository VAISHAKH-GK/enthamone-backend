package handlers

import (
	"github.com/FulgurCode/enthamone-backend/pkg/message"
	"github.com/FulgurCode/enthamone-backend/pkg/ws"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

// Handle websocket connection
func HandleConnection(c *websocket.Conn) {
	// Generate id and send to client
	var id string = uuid.NewString()
	var msg = message.Message{
		To:          id,
		MessageType: message.ID,
		Content:     id,
	}
	var err = c.WriteJSON(msg)
	if err == websocket.ErrCloseSent {
		return
	}

	// Create new client
	var client = ws.NewClient(c,id)
	go client.ListenMsg()
	client.WriteMsg()
}
