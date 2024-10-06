package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// NOTE: maintain tight coupling here unlike other APIs with handler - service - repository pattern
// since this logic is tighly linked with the primary websocket server of the application.

/**
* Upgrades connection to websocket connection
**/
func (s *Server) HandleWebSocket(c *gin.Context) {
	// upgrade the HTTP connection to a WebSocket connection
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upgrade connection"})
		return
	}
	defer conn.Close()

	s.serverChan <- GameMove{
		Action:  "join_match",
		Payload: "123",
	}

	// set conenction and player
	s.clientConns[conn] = Player{id: "1", name: "guest"}

	// handle each connected client's messages concurrently
	go s.ServeWsClient(conn)
}

/**
* Serves each individual connected client
**/
func (s *Server) ServeWsClient(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println("Received message:", message)

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}

/**
* Websocket Message Hub to handle all messages.
**/
func (s *Server) MessageHub() {
	for {
		select {
		case gameMove := <-s.serverChan:
			fmt.Printf("Game move received: %+v\n\n", gameMove)
		}
	}
}
