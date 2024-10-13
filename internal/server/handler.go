package server

import (
	"encoding/json"
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
func (s *MultiplayerServer) HandleMatchConn(c *gin.Context) {
	// upgrade the HTTP connection to a WebSocket connection
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println("Error establishing websocket connection.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upgrade connection"})
		return
	}

	// handle each connected client's messages concurrently
	go s.ServeConnectedPlayer(conn)
}

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

		// decode message to pre-defined json structure
		var decodedMsg GameMessage

		err = json.Unmarshal(message, &decodedMsg)

		if err != nil {
			fmt.Println("Error when decoding payload.")
			conn.WriteJSON(GameMessage{Action: "Error", Payload: "Your message to server was the incorrect format and could not be decoded as JSON."})
		}

		clientPackage := ClientPackage{GameMessage: decodedMsg, Conn: conn}

		// send message to MessageHub for handling based on type
		s.serverChan <- clientPackage
	}
}

/**
* Adds a player to the list of client connections.
**/
func (s *MultiplayerServer) addClient(conn *websocket.Conn, client Player) {
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

/**
* Stops a match and removes it from the map of matches that a player belongs to. TODO: stop the match
**/
func (s *MultiplayerServer) StopMatch(playerId uuid.UUID) error {
	// search for player in a match
	for matchIndex, players := range s.matches {
		for _, player := range players {
			if player.ID == playerId {
				// stop and remove match
				delete(s.matches, matchIndex)
				return nil
			}
		}
	}
	return fmt.Errorf("No player with this id was found in any on-going match.")
}

/**
* Helps find a match for the player.
**/
func (s *MultiplayerServer) findMatch(player Player) uuid.UUID {
	fmt.Println("Finding match...")

	// loop through current matches and find an opponent still waiting
	for matchId, match := range s.matches {
		// check length of match to know if its full
		var matchFull bool = false
		fmt.Println("Length of match:", len(match))

		// match is full is length of match has reached 2
		matchFull = len(match) == 2

		// join match if not full
		if !matchFull {
			s.matches[matchId] = append(s.matches[matchId], player)
			// end search
			return matchId
		}
	}

	// iteration over, meaning all matches are full, create a new one

	newMatch := []Player{player}

	newMatchUuid := uuid.New()
	s.matches[newMatchUuid] = newMatch

	return newMatchUuid
}

type PlayerIdString struct {
	id   string
	name string
}

// For pretty-fying matches for easier testing by mapping each id from a UUID
// to a string
func MapIdStringMatches(matches map[uuid.UUID][]Player) map[string][]PlayerIdString {
	matchesToPrint := make(map[string][]PlayerIdString)

	// map over and convert byte slice keys to id strings
	for index := range matches {
		var player1, player2 PlayerIdString

		if len(matches[index]) > 0 {
			player1 = PlayerIdString{
				id:   matches[index][0].ID.String(),
				name: matches[index][0].Name,
			}
		}

		if len(matches[index]) > 1 {
			player2 = PlayerIdString{
				id:   matches[index][1].ID.String(),
				name: matches[index][1].Name,
			}
		}

		matchesToPrint[index.String()] = []PlayerIdString{player1, player2}
	}

	// print result
	fmt.Printf("PRETTY PRINT MATCHES: %v\n\n\n", matchesToPrint)

	return matchesToPrint
}
