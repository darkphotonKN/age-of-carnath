package server

import (
	"fmt"

	"github.com/darkphotonKN/age-of-carnath/internal/game"
	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

/**
* Wrapper function for goroutine to find a match and respond the client(s) with the game state.
**/
func (s *MultiplayerServer) findMatchAndBroadcast(p models.Player) {
	// handles communication match find results
	matchFindChan := make(chan uuid.UUID)

	// find the match concurrently, sending to the channel once it's done
	go func() {
		id := s.findMatch(p)
		matchFindChan <- id
	}()

	// wait until match is found
	id := <-matchFindChan

	// broadcast to all users participating in the match
	s.broadcastGameStateToPlayers(id)
}

/**
* All Match-making Business Logic
**/

/**
* Helps find a match for the player.
* TODO: For v1.1: Add matchmaking algorithm.
**/
func (s *MultiplayerServer) findMatch(player models.Player) uuid.UUID {
	// maps are not thread-safe, adding locking incase match was removed / altered
	s.mu.Lock()
	defer s.mu.Unlock()

	// loop through current matches and find an opponent still waiting
	for matchId, game := range s.matches {

		match := game.Players

		// check length of match to know if its full
		var matchFull bool = false
		fmt.Println("Length of match:", len(match))

		// match is "full" is when the length of match has reached 2
		matchFull = len(match) == 2

		// join match if not full
		if !matchFull {
			game.JoinGame(&player)

			// end search
			return matchId
		}
	}

	// iteration over, meaning all matches are full, create a new one

	// initalize a game
	newGame := game.InitializeGame(&player)

	s.matches[newGame.ID] = newGame

	return newGame.ID
}

/**
* Broadcasts current game state to all players of a particular match.
**/
func (s *MultiplayerServer) broadcastGameStateToPlayers(matchId uuid.UUID) {
	// TODO: Refactor to include better way of access player connections.
	gameState := s.matches[matchId]

	// loop through all players of the game and find corresponding
	// client's websocket connection to broadcast
	// NOTE: If multiple goroutines try to write to the same connection at the same time,
	// this can cause data races or undefined behavior.

	for _, player := range gameState.Players {
		for conn, client := range s.clientConns {
			if player.ID == client.ID {
				// get current channel responsible for reading-in messages before writing back to client
				msgChan := s.getGameMsgChan(conn)
				msgForClient := GameMessage{
					Action:  init_match,
					Payload: *gameState,
				}
				msgChan <- msgForClient
			}
		}
	}
}

/**
* Handles adding clients and creating gameMsgChans for handling connection writes
* back to the connected client.
**/
func (s *MultiplayerServer) setupClientWriter(conn *websocket.Conn) {
	err := s.createGameMsgChan(conn)

	if err != nil {
		fmt.Printf("Error when attempting to creating client writer: %s\n", err)
		return
	}

	// in the case the channel exists
	if msgChan := s.getGameMsgChan(conn); msgChan != nil {

		// concurrently listen to all incoming messages over the channel to write back to client
		go func() {
			for msg := range msgChan {
				err := conn.WriteJSON(msg)

				if err != nil {
					// TODO: remove connection from channel and close
					s.cleanUpClient(conn)
					break
				}
			}
		}()
	}
}

/**
* Cleans up both clients and client writer.
**/
func (s *MultiplayerServer) cleanUpClient(conn *websocket.Conn) {
	// NOTE: keep mutex due to multiple concurrent access to global resources
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.clientConns[conn]

	// end prematurely due to client already removed / missing
	if !ok {
		return
	}

	// close channel
	msgChan := s.getGameMsgChan(conn)
	close(msgChan)

	// remove msgChan from list of game message channels
	delete(s.gameMessageChans, conn)

	// remove from list of clients
	delete(s.clientConns, conn)
}

// --- Helpers ---

type PlayerIdString struct {
	id   string
	name string
}

// For pretty-fying matches for easier testing by mapping each id from a UUID
// to a string
func MapIdStringMatches(matches map[uuid.UUID]*game.Game) map[string][]PlayerIdString {
	matchesToPrint := make(map[string][]PlayerIdString)

	// map over and convert byte slice keys to id strings
	for index := range matches {
		var player1, player2 PlayerIdString

		if len(matches[index].Players) > 0 {
			player1 = PlayerIdString{
				id:   matches[index].Players[0].ID.String(),
				name: matches[index].Players[0].Name,
			}
		}

		if len(matches[index].Players) > 1 {
			player2 = PlayerIdString{
				id:   matches[index].Players[1].ID.String(),
				name: matches[index].Players[1].Name,
			}
		}

		matchesToPrint[index.String()] = []PlayerIdString{player1, player2}
	}

	// print result
	fmt.Printf("PRETTY PRINT MATCHES: %v\n\n\n", matchesToPrint)

	return matchesToPrint
}
