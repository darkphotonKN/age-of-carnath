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

	// add player with a unique id to list of connections with their unique ws connection
	// as a key
	newPlayerId := uuid.New()
	s.clientConns[conn] = Player{id: newPlayerId, name: "guest"}

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

		// decode message to pre-defined json structure
		var decodedMsg GameMessage

		err = json.Unmarshal(message, &decodedMsg)

		if err != nil {
			fmt.Println("Error when decoding payload.")
			conn.WriteJSON(GameMessage{Action: "Error", Payload: "Your message to server was the incorrect format and could not be decoded as JSON."})
		}

		fmt.Println("Received message as string:", string(message))

		// send message to MessageHub for handling based on type
		s.serverChan <- decodedMsg

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
		case gameMessage := <-s.serverChan:
			fmt.Printf("Game message received: %+v\n\n", gameMessage)
			switch gameMessage.Action {
			case "find_match":
				// TODO: update this to be their actual player from payload
				s.findMatch(Player{id: uuid.New(), name: "Second player"})
			}
		}
	}
}

/**
* Helps find a match for the player.
**/
func (s *MultiplayerServer) findMatch(player Player) uuid.UUID {
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
	newMatch := make([]Player, 2)
	newMatch[0] = player

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
				id:   matches[index][0].id.String(),
				name: matches[index][0].name,
			}
		}

		if len(matches[index]) > 1 {
			player2 = PlayerIdString{
				id:   matches[index][1].id.String(),
				name: matches[index][1].name,
			}
		}

		matchesToPrint[index.String()] = []PlayerIdString{player1, player2}
	}

	// print result
	fmt.Printf("PRETTY PRINT MATCHES: %v\n\n\n", matchesToPrint)

	return matchesToPrint
}
