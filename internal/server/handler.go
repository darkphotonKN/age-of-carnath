package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// -- Core Handler --

// NOTE: maintain tight coupling here unlike other APIs with handler - service - repository pattern
// since this logic is tighly linked with the primary websocket server of the application.

/**
* Handles player searching for a match (incoming connection), upgrades them to websocket connections,
* and passes them off to individual goroutines to be handled concurrently.
**/
func (s *MultiplayerServer) HandleMatchConn(c *gin.Context) {
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println("Error establishing websocket connection.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upgrade connection"})
		return
	}

	// handle each connected client's messages concurrently
	go s.ServeConnectedPlayer(conn)
}

// -- Primary Methods --

/**
* Serves each individual connected player.
* NOTE: Gorilla Websocket package only allows ONE CONCURRENT WRITER
* at a time, meaning its best to utilize *unbuffered* channels to prevent
* a single client from locking the entire server.
**/

func (s *MultiplayerServer) ServeConnectedPlayer(conn *websocket.Conn) {
	// removes client and closes connection
	defer func() {
		fmt.Println("Connection closed due to end of function.")
		s.removeClient(conn)
	}()

	fmt.Printf("Starting listener for user %v\n", s.clientConns[conn])
	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			// -- clean up connection --

			// handle error - if error is of type thats an unknown error
			// that matches the two types listed, we close return the loop and
			// close it immediately (via the defer)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("Abormal error occured with player %v. Closing connection.\n", s.clientConns[conn])
			}
			break
		}

		// decode message to pre-defined json structure "GameMessage"
		var decodedMsg GameMessage

		err = json.Unmarshal(message, &decodedMsg)

		if err != nil {
			fmt.Println("Error when decoding payload.")
			conn.WriteJSON(GameMessage{Action: "Error", Payload: "Your message to server was the incorrect format and could not be decoded as JSON."})
			continue
		}

		clientPackage := ClientPackage{GameMessage: decodedMsg, Conn: conn}

		// Send message to MessageHub via an *unbuffered channel* for handling based on type.
		s.serverChan <- clientPackage
	}
}

/**
* Adds a player to the list of client connections.
**/
func (s *MultiplayerServer) addClient(conn *websocket.Conn, client models.Player) {
	// lock and unlock to prevent race conditions
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clientConns[conn] = client
}

/**
* Removes a player from the list of client connections.
**/
func (s *MultiplayerServer) removeClient(conn *websocket.Conn) {
	// lock and unlock to prevent race conditions
	s.mu.Lock()
	defer s.mu.Unlock()

	if client, ok := s.clientConns[conn]; ok {
		conn.Close()

		// remove them from the match and stops the match
		s.StopMatch(client.ID)

		// remove from list of connections
		delete(s.clientConns, conn)
	}
}

/*
*
* Stops a match and removes it from the map of matches that a player belongs to.
* TODO: stop the match for all players.
*
 */
func (s *MultiplayerServer) StopMatch(playerId uuid.UUID) error {
	// search for player in a match
	for matchIndex, game := range s.matches {
		for _, player := range game.Players {
			if player.ID == playerId {
				// stop and remove match
				delete(s.matches, matchIndex)
				return nil
			}
		}
	}
	return fmt.Errorf("No player with this id was found in any on-going match.")
}
