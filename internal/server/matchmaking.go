package server

import (
	"fmt"

	"github.com/darkphotonKN/age-of-carnath/internal/game"
	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
)

/**
* All Match-making Business Logic
**/

/**
* Helps find a match for the player.
* TODO: For v1.1: Add matchmaking algorithm.
**/
func (s *MultiplayerServer) findMatch(player models.Player) uuid.UUID {
	// maps are not thread-safe, can add locking to be sure incase match was removed / altered
	s.mu.Lock()
	defer s.mu.Unlock()

	// loop through current matches and find an opponent still waiting
	for matchId, game := range s.matches {

		match := game.Players

		// check length of match to know if its full
		var matchFull bool = false
		fmt.Println("Length of match:", len(match))

		// match is "full" is length of match has reached 2
		matchFull = len(match) == 2

		// join match if not full
		if !matchFull {

			// s.matches[matchId].Players = append(s.matches[matchId].Players, player)
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
* Wrapper function for goroutine to find a match and respond the client(s) with the game state.
**/
func (s *MultiplayerServer) FindMatchAndBroadcast(p models.Player) {
	// find the match concurrently
	go s.findMatch(p)

	// wait until match is found
	// TODO:
	// - Create channel to communicate match found concurrently.
	// - Update findMatch() to return multiple things, like match information not just the id.
	// - Update tests to follow new flow (with channels).
	// game := <-matchFound

	// broadcast to all users participating in the match

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
