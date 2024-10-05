package game

import (
	"fmt"
	"testing"

	"github.com/darkphotonKN/age-of-carnath/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestNewGame_GridState(t *testing.T) {
	server := server.NewServer(":3333")
	game := NewGame(server, 2, 2)

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
		}, {
			{
				position: Position{x: 0, y: 1},
				content:  Content{contentType: empty},
			},

			{
				position: Position{x: 1, y: 1},
				content:  Content{contentType: empty},
			},
		},
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("game.GridState", game.GridState)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("expectedGridState", expectedGridState)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	assert.Equal(t, expectedGridState, game.GridState)
}
