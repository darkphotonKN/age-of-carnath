package server

import (
	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
)

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

/**
* Broadcasts a error to all players in a match.
**/
func (s *MultiplayerServer) broadcastGameErrorToPlayers(matchId uuid.UUID, errorMessage string) {
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
					Action:  match_error,
					Payload: errorMessage,
				}

				msgChan <- msgForClient
			}
		}
	}
}

/**
* Broadcasts a single message to a single player.
**/

func (s *MultiplayerServer) sendMessageToPlayer(player models.Player, action Action, msg string) {

	for conn, client := range s.clientConns {
		if client.ID == player.ID {

			// send message to the single player
			msgChan := s.getGameMsgChan(conn)

			msgForClient := GameMessage{
				Action:  action,
				Payload: msg,
			}

			msgChan <- msgForClient
		}
	}

}
