package server

import (
	"fmt"

	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
)

/**
* -- Game Message --
* The payload for communication between client-server.
*
* Handles all the conversion and parsing for the game messages sent from client to
* the server.
**/

type Action string

const (
	find_match  Action = "find_match"
	init_match  Action = "init_match"
	match_error Action = "match_error"
	move        Action = "move"
	attack      Action = "attack"
)

type GameMessage struct {
	Action  Action      `json:"action"`
	Payload interface{} `json:"payload"`
}

/**
* Provides conversion of the Payload based on the action type.
**/
func (gm *GameMessage) ParsePayload() error {
	switch gm.Action {
	case find_match:

		// assume payload is Player
		payloadMap, ok := gm.Payload.(map[string]interface{})
		if !ok {
			fmt.Println("Payload is not of type map[string]interface{}")
			return fmt.Errorf("Payload is not of type map[string]interface{}")
		}

		fmt.Printf("Converting to player: %+v\n", payloadMap)

		// Extract the "id" and "name" fields from the map
		idStr, idOk := payloadMap["id"].(string)
		name, nameOk := payloadMap["name"].(string)

		if !idOk || !nameOk {
			fmt.Println("Error: The required fields 'id' or 'name' are not in the expected format.")
			return fmt.Errorf("Error: The required fields 'id' or 'name' are not in the expected format.")
		}

		fmt.Println("idStr:", idStr)
		fmt.Println("name:", name)

		idUUID, err := uuid.Parse(idStr)

		if err != nil {
			fmt.Println("Could not parse id into a UUID.")
			return fmt.Errorf("Could not parse id into a UUID.")
		}

		convertedPlayer := models.Player{
			ID:   idUUID,
			Name: name,
		}

		fmt.Println("Converted player:", convertedPlayer)

		gm.Payload = convertedPlayer

	default:
		return fmt.Errorf("No valid Action from game message.")
	}

	return nil
}
