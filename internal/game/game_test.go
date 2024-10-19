package game

import (
	"fmt"
	"testing"

	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_NewGame_GridState(t *testing.T) {
	game := NewGame(uuid.New(), 3, 3)

	expectedGridState := GridState{
		{
			{
				ContentType: empty,
				Position:    Position{x: 0, y: 0},
			},
			{
				ContentType: empty,
				Position:    Position{x: 1, y: 0},
			},
			{
				ContentType: empty,
				Position:    Position{x: 2, y: 0},
			},
		},
		{
			{
				ContentType: empty,
				Position:    Position{x: 0, y: 1},
			},

			{
				ContentType: empty,
				Position:    Position{x: 1, y: 1},
			},
			{
				ContentType: empty,
				Position:    Position{x: 2, y: 1},
			},
		},
		{
			{
				ContentType: empty,
				Position:    Position{x: 0, y: 2},
			},
			{
				ContentType: empty,
				Position:    Position{x: 1, y: 2},
			},
			{
				ContentType: empty,
				Position:    Position{x: 2, y: 2},
			},
		},
	}

	assert.Equal(t, expectedGridState, game.GridState)
}

// Test player that created the match spawned correctly.
func Test_SpawnPlayerOnGrid(t *testing.T) {

	// create mock player
	mockPlayer := models.Player{
		ID:   uuid.New(),
		Name: "mock player",
	}

	game := NewGame(uuid.New(), 3, 3)

	// spawns player in grid at a random location
	game.SpawnPlayerOnGrid(&mockPlayer)

	// test if player exists at all
	playerExists := false
	for _, row := range game.GridState {

		for _, block := range row {
			// can check for nil, as its a pointer
			if block.Content.Player != nil && block.Content.Player.ID == mockPlayer.ID {

				fmt.Printf("PLAYER EXISTS!!! Player: %+v\n\n", block)
				playerExists = true
			}
		}
	}

	if !playerExists {
		fmt.Printf("gridState: %+v\n\n", game.GridState)
		t.Errorf("Expected player to be spawned, but player did not exist inside the grid state.")
	}
}
