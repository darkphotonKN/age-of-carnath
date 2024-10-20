package server

import (
	"fmt"

	"github.com/darkphotonKN/age-of-carnath/internal/game"
	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
)

/**
* Websocket Message Hub to handle all messages.
**/
func (s *MultiplayerServer) MessageHub() {
	fmt.Println("Starting Message Hub")

	for {
		fmt.Printf("Current client connections in session: %+v\n\n", s.clientConns)
		printOngoingMatches(s.matches)

		select {
		case clientPackage := <-s.serverChan:
			fmt.Printf("Client Package received: %+v\n\n", clientPackage)

			// NOTE: parses payload to a specific type based on the action type
			// e.g. when payload is "find_match" the payload is converted from interface{} -> Player
			err := clientPackage.GameMessage.ParsePayload()

			if err != nil {
				fmt.Printf("Error occured when attempting to parse payload: %s\n", err)
				clientPackage.Conn.WriteJSON("Error attempting to parse payload.")
				continue
			}

			switch clientPackage.GameMessage.Action {
			case "find_match":
				fmt.Println("Inside 'find match' case, payload:", clientPackage.GameMessage.Payload)

				// assert Payload type specific to gameMessage.Action == "find_match", which is Player
				player, ok := clientPackage.GameMessage.Payload.(models.Player)

				if !ok {
					fmt.Printf("Error attempting to assert player from payload.\n")
					clientPackage.Conn.WriteJSON("Error attempting to assert player from payload.")
					continue
				}

				s.addClient(clientPackage.Conn, player)

				// initiating finding a match for the player
				s.findMatch(player)
			}
		}
	}
}

func printOngoingMatches(matches map[uuid.UUID]*game.Game) {
	fmt.Println("Current Matches")
	fmt.Println("---------------")

	for id, match := range matches {

		// filter out empty cells for testing
		var nonEmptyGrid []game.GridBlock

		for _, row := range match.GridState {
			for _, block := range row {
				if block.ContentType != game.EmptyType {
					nonEmptyGrid = append(nonEmptyGrid, block)
				}
			}
		}

		fmt.Printf("Match %s\nplayers: %+v\nNon-Empty GridState: %+v\n\n", id, match.Players, nonEmptyGrid)
	}
}
