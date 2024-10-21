package game

import (
	"fmt"
	"testing"

	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_NewGame_GridState(t *testing.T) {
	game := NewGame(3, 3)

	expectedGridState := GridState{
		{
			{
				ContentType: EmptyType,
				Position:    Position{x: 0, y: 0},
			},
			{
				ContentType: EmptyType,
				Position:    Position{x: 1, y: 0},
			},
			{
				ContentType: EmptyType,
				Position:    Position{x: 2, y: 0},
			},
		},
		{
			{
				ContentType: EmptyType,
				Position:    Position{x: 0, y: 1},
			},

			{
				ContentType: EmptyType,
				Position:    Position{x: 1, y: 1},
			},
			{
				ContentType: EmptyType,
				Position:    Position{x: 2, y: 1},
			},
		},
		{
			{
				ContentType: EmptyType,
				Position:    Position{x: 0, y: 2},
			},
			{
				ContentType: EmptyType,
				Position:    Position{x: 1, y: 2},
			},
			{
				ContentType: EmptyType,
				Position:    Position{x: 2, y: 2},
			},
		},
	}

	assert.Equal(t, expectedGridState, game.GridState)
}

// Test player that created the match spawned correctly.
func Test_SpawnPlayerOnGrid(t *testing.T) {

	// create mock player
	playerId := uuid.New()
	mockPlayer := models.Player{
		ID:   playerId,
		Name: "mock player",
	}

	game := NewGame(3, 3)

	// spawns player in grid at a random location
	game.SpawnPlayerOnGrid(&mockPlayer)

	// test if player exists at all
	playerExists := false

	for _, row := range game.GridState {
		for _, block := range row {
			// can check for nil, as its a pointer
			if block.Content.Player != nil && block.Content.Player.ID == playerId {

				fmt.Printf("PLAYER EXISTS!!! Player: %+v\n\n", block)
				playerExists = true
			}
		}
	}

	if !playerExists {
		fmt.Printf("gridState: %+v\n\n", game.GridState)
		t.Errorf("Expected player to be spawned, but player did not exist inside the grid state.")
	}

	// Test Players Slice
	playerExists = false // reset from first test
	for _, player := range game.Players {
		if player.ID == playerId {
			playerExists = true
		}
	}

	if !playerExists {
		fmt.Printf("gridState: %+v\n\n", game.GridState)
		t.Errorf("Player did not found in player slice.")
	}
}

// Test PLayer joining a pre-existing match.
func Test_JoinGame(t *testing.T) {
	game := NewGame(2, 2)

	// create mock player
	playerId := uuid.New()

	mockPlayer := models.Player{
		ID:   playerId,
		Name: "Mock Player 1",
	}

	game.JoinGame(&mockPlayer)

	// search grid that player exists
	playerExists := false

	for _, row := range game.GridState {
		for _, block := range row {
			// can check for nil, as its a pointer
			if block.Content.Player != nil && block.Content.Player.ID == playerId {

				fmt.Printf("PLAYER EXISTS!!! Player: %+v\n\n", block)
				playerExists = true
			}
		}
	}

	if !playerExists {
		fmt.Printf("gridState: %+v\n\n", game.GridState)
		t.Errorf("Expected player to have joined game, but is but player did not exist inside the grid state.")
	}

	// Test Players Slice

	// check player exists in the player slice
	playerExists = false // reset from first test
	for _, player := range game.Players {
		if player.ID == playerId {
			playerExists = true
		}
	}

	if !playerExists {
		fmt.Printf("gridState: %+v\n\n", game.GridState)
		t.Errorf("Player did not found in player slice.")
	}
}
