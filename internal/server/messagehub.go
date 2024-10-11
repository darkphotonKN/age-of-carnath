package server

import (
	"fmt"
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
		case clientPackage := <-s.serverChan:
			fmt.Printf("Client Package received: %+v\n\n", clientPackage)

			gameMessage := clientPackage.GameMessage

			switch gameMessage.Action {
			case "find_match":
				// add player with a unique id to list of connections with their unique ws connection
				// as a key
				player, ok := gameMessage.Payload.(Player) // assert that player was the payload in the case of find match
				if !ok {
					clientPackage.Conn.WriteJSON("Player was not in the payload of join_match action.")
					continue
				}
				s.addClient(clientPackage.Conn, player)

				// initiating finding a match for the player
				s.findMatch(player)
			}

		}
	}
}
