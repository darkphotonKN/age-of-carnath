package server

import (
	"fmt"

	"github.com/google/uuid"
)

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

				fmt.Println("Finding a match...")
				s.findMatch(Player{id: uuid.New(), name: "Second player"})
			}
		}
	}
}
