package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// NOTE: maintain tight coupling here unlike other APIs with handler - service - repository pattern
// since this logic is tighly linked with the primary websocket server of the application.

/**
* Handles player searching for a match (incoming connection), upgrades them to webscoket connections,
* and passes them off to individual goroutines to be handled concurrently.
**/
func (s *MultiplayerServer) HandleMatchSearch(c *gin.Context) {
	// upgrade the HTTP connection to a WebSocket connection
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println("Error establishing websocket connection.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upgrade connection"})
		return
	}

	s.serverChan <- GameMove{
		Action:  "join_match",
		Payload: "123",
	}

	// set connection and player
	s.clientConns[conn] = Player{id: "1", name: "guest"}

	// find match for player
	s.findMatch(conn)

	// handle each connected client's messages concurrently
	go s.ServeConnectedPlayer(conn)
}

/**
* Serves each individual connected player
**/
func (s *MultiplayerServer) ServeConnectedPlayer(conn *websocket.Conn) {
	defer func() {
		fmt.Println("Connection closed due to end of function.")
		conn.Close()
	}()

	fmt.Printf("Starting listener for user %v", s.clientConns[conn])
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
func (s *MultiplayerServer) MessageHub() {
	fmt.Println("Starting Message Hub")

	for {
		fmt.Printf("Current client connections in session: %+v\n\n", s.clientConns)
		fmt.Printf("Current ongoing matches %+v\n\n", s.matches)
		select {
		case gameMove := <-s.serverChan:
			fmt.Printf("Game move received: %+v\n\n", gameMove)
		}
	}
}

/**
* Helps find a match for the player.
**/
func (s *MultiplayerServer) findMatch(conn *websocket.Conn) {
	// loop through current matches and find an opponent still waiting
	for matchId, match := range s.matches {
		// check length of match to know if its full
		var matchFull bool = false
		if len(match) == 2 {
			matchFull = true
		}

		// join match if not full
		if !matchFull {
			s.matches[matchId] = append(s.matches[matchId], Player{id: "123", name: "placeholder"})
		}
	}

	// iteration over, meaning all matches are full, create a new one
	newMatch := make([]Player, 2)
	newMatch = append(newMatch, Player{id: "321", name: "creatorplaceholder"})

	// TODO: generate a new UUID
	newMatcUuid := uuid.New()
	s.matches[newMatcUuid] = newMatch
}
