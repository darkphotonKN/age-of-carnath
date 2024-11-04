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
func (s *MultiplayerServer) findMatchAndBroadcast(player models.Player) {
	// handles communication match find results
	matchFindChan := make(chan uuid.UUID)

	go func() {
		id := s.findMatch(player)

		fmt.Println("Match found initially, id:", id)

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
			fmt.Println("Time passed (seconds):", secondsPassed)

		// match found so broadcast info to all users participating in the match
		case matchFoundId := <-matchFindChan:
			fmt.Println("Game found.")
			s.broadcastGameStateToPlayers(matchFoundId)
			return

			// timeout passed first so stop the match find, send error
		case <-timeout:
			s.sendMessageToPlayer(player, match_error, "Timeout when searching for match.")
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
* 3) Add mutex locks.
* 4) For v1.1: Add matchmaking algorithm.
**/
func (s *MultiplayerServer) findMatch(player models.Player) uuid.UUID {
	fmt.Println("Starting matchmaking..")

	// loop through current matches and find an opponent still waiting
	// keep looping until queue is over 2 minutes long
	// search all matches for empty match
	for matchId, game := range s.matches {
		// maps are not thread-safe, adding locking incase match was removed / altered
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

	fmt.Println("No matches, starting new one..")

	// after first iteration means there were no open games, create one
	newMatch := game.InitializeGame(&player)

	var timeWaited time.Duration = time.Second * 0
	waitTime := time.Second * 1

	for {
		if timeWaited >= time.Second*9 {
			return uuid.Nil
		}

		// repeat check for game start every 3 seconds
		timeWaited += waitTime
		time.Sleep(waitTime)

		match, ok := s.matches[newMatch.ID]
		fmt.Printf("Time Waited: %d\nMatch Check: %+v\nMatch Found: %+v\n", timeWaited, s.matches[newMatch.ID], ok)

		if !ok {
			continue
		}

		fmt.Printf("match length currently: %d", len(match.Players))

		// match full, return id and start match
		if len(match.Players) == 2 {
			return match.ID
		}
	}

}

// --- Helpers ---

type PlayerIdString struct {
	id   string
	name string
}

/**
* For pretty-fying matches for easier testing by mapping each id from a UUID
* to a string.
**/
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
