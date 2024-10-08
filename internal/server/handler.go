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
* Helps find a match for the player.
**/
func (s *MultiplayerServer) findMatch(player Player) {
	// loop through current matches and find an opponent still waiting
	for matchId, match := range s.matches {
		// check length of match to know if its full
		var matchFull bool = false
		if len(match) == 2 {
			matchFull = true
		}

		// join match if not full
		if !matchFull {
			s.matches[matchId] = append(s.matches[matchId], player)
		}
	}

	// iteration over, meaning all matches are full, create a new one
	newMatch := make([]Player, 2)
	newMatch = append(newMatch, player)

	// TODO: generate a new UUID
	newMatchUuid := uuid.New()
	s.matches[newMatchUuid] = newMatch
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

type PrettyPrintPlayer struct {
	id   string
	name string
}

func PrettyPrintMatches(matches map[uuid.UUID][]Player) {
	matchesToPrint := make(map[string][]PrettyPrintPlayer)

	// map over and convert byte slice keys to id strings
	for index := range matches {
		player1 := PrettyPrintPlayer{
			id:   matches[index][0].id.String(),
			name: matches[index][0].name,
		}

		player2 := PrettyPrintPlayer{
			id:   matches[index][1].id.String(),
			name: matches[index][1].name,
		}

		matchesToPrint[index.String()] = []PrettyPrintPlayer{player1, player2}
	}

	// print result
	fmt.Printf("PRETTY PRINT MATCHES: %+v\n\n\n", matchesToPrint)
}
