package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame_GridState(t *testing.T) {
	game := NewGame(3, 3)

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
