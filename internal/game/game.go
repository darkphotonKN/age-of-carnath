package game

import (
	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
)

/**
* Holds all grid and game information for a **single** match.
**/
type Game struct {
	// Game unique identifier. NOTE: Currently matches the map key for server matches.
	ID uuid.UUID

	// holds all the game's grid information
	GridState GridState

	// contains the match's players. NOTE: (max length 2)
	Players []models.Player
}

/**
* Position on the Grid
**/
type Position struct {
	x uint8
	y uint8
}

/**
* Enum of Content Types
**/
type ContentType string

const (
	empty  ContentType = "empty"
	player ContentType = "player"
	item   ContentType = "item"
)

/**
* The content that can occupy a position on the GridState.
**/

type Item struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

type Content struct {
	Player *models.Player `json:"player,omitempty"`
	Item   *Item          `json:"item,omitempty"`
}

/**
* Represents a Single Grid Block
* Position - the occupying position of this block in relation to the entire grid.
* Content - the occupying content of the grid.
**/
type GridBlock struct {
	Position    Position    `json:"position"`
	ContentType ContentType `json:"contentType"`
	Content     Content     `json:"content"`
}

/**
* Represents the entire game grid, compopsed of GridBlocks.
**/
type GridState [][]GridBlock

func NewGame(gridRows uint8, gridCols uint8) *Game {
	newGrid := initializeGrid(gridRows, gridCols)

	return &Game{
		GridState: newGrid,
	}
}

/**
* Initalizes base game grid
**/
func initializeGrid(rows uint8, cols uint8) GridState {
	newGridState := make([][]GridBlock, rows)

	for rowIndex := range newGridState {
		newCols := make([]GridBlock, cols)
		newGridState[rowIndex] = newCols

		for colIndex := range newGridState[rowIndex] {
			newGridState[rowIndex][colIndex] = GridBlock{
				Position: Position{
					x: uint8(colIndex),
					y: uint8(rowIndex),
				},
				ContentType: empty,
			}
		}
	}
	return newGridState
}
