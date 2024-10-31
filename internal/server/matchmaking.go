package server

import (
	"fmt"
	"time"

	"github.com/darkphotonKN/age-of-carnath/internal/game"
	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
)

/**
* Wrapper function for goroutine to find a match and respond the client(s) with the game state.
**/
func (s *MultiplayerServer) findMatchAndBroadcast(p models.Player) {
	// handles communication match find results
	matchFindChan := make(chan uuid.UUID)

	go func() {
		id := s.findMatch(p)

		matchFindChan <- id
	}()

	// for tracking how long player has waited in queue
	ticker := time.NewTicker(time.Second * 1) // counts at 1 second per interval
	timeout := time.After(time.Second * 10)
	secondsPassed := 0

	for {
		select {
		// case that a ticker interval passed, increment and check on the ticker channel
		case <-ticker.C:
			secondsPassed += 1
			fmt.Println("Seconds passed:", secondsPassed)

		// match found so broadcast info to all users participating in the match
		case matchFoundId := <-matchFindChan:
			s.broadcastGameStateToPlayers(matchFoundId)

			// timeout passed first so stop the matchfind
		case <-timeout:
			return
		}
	}
}

/**
* All Match-making Business Logic
**/

/**
* Helps find a match for the player.
*
* TODO:
* 1) Fix close error for client (1001 going away). DONE
* 2) Only allow init match once the game is full, otherwise matchmaking should be pending.
* 3) For v1.1: Add matchmaking algorithm.
**/
func (s *MultiplayerServer) findMatch(player models.Player) uuid.UUID {
	// maps are not thread-safe, adding locking incase match was removed / altered
	s.mu.Lock()
	defer s.mu.Unlock()

	// loop through current matches and find an opponent still waiting
	// keep looping until queue is over 2 minutes long
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

			// end search immediately
			return matchId
		}
	}

	// iteration over, meaning all matches are full, create a new one

	// initalize a game
	newGame := game.InitializeGame(&player)

	s.matches[newGame.ID] = newGame

	return newGame.ID

	// // but only return once the game has been "filled up" (2 players)
	// for {
	// 	time.Sleep(time.Second * 2000)
	//
	// 	fmt.Printf("No of Players: %d\nMatch Players: %+v\n", len(s.matches[newGame.ID].Players), s.matches[newGame.ID].Players)
	//
	// 	if len(s.matches[newGame.ID].Players) == 2 {
	// 		return newGame.ID
	// 	}
	// }

}

/**
* Broadcasts current game state to all players of a particular match.
**/
func (s *MultiplayerServer) broadcastGameStateToPlayers(matchId uuid.UUID) {
	// TODO: Refactor MultiplayerServer struct to include info
	// for simpler way of accessing player connections.
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
