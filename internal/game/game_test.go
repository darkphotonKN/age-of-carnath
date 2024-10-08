package game

import (
	"testing"

	"github.com/darkphotonKN/age-of-carnath/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestNewGame_GridState(t *testing.T) {
	server := server.NewMultiplayerServer(":3333")
	game := NewGame(server, 3, 3)

	expectedGridState := GridState{
		{
			{
				position: Position{x: 0, y: 0},
				content:  Content{contentType: empty},
			},
			{
				position: Position{x: 1, y: 0},
				content:  Content{contentType: empty},
			},
			{
				position: Position{x: 2, y: 0},
				content:  Content{contentType: empty},
			},
		},
		{
			{
				position: Position{x: 0, y: 1},
				content:  Content{contentType: empty},
			},

			{
				position: Position{x: 1, y: 1},
				content:  Content{contentType: empty},
			},
			{
				position: Position{x: 2, y: 1},
				content:  Content{contentType: empty},
			},
		},
		{
			{
				position: Position{x: 0, y: 2},
				content:  Content{contentType: empty},
			},
			{
				position: Position{x: 1, y: 2},
				content:  Content{contentType: empty},
			},
			{
				position: Position{x: 2, y: 2},
				content:  Content{contentType: empty},
			},
		},
	}

	assert.Equal(t, expectedGridState, game.GridState)
}
