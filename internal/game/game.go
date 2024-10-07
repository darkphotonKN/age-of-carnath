package game

import "github.com/darkphotonKN/age-of-carnath/internal/server"

/**
* Holds all grid and game information.
* Uses DI for server access (via pointer).
**/
type Game struct {
	server    *server.MultiplayerServer
	GridState GridState
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
type Content struct {
	contentType ContentType
	value       interface{}
}

/**
* Represents a Single Grid Block
* Position - the occupying position of this block in relation to the entire grid.
* Content - the occupying content of the grid.
**/
type GridBlock struct {
	position Position
	content  Content
}

/**
* Represents the entire game grid, compopsed of GridBlocks.
**/
type GridState [][]GridBlock

func NewGame(server *server.MultiplayerServer, gridRows uint8, gridCols uint8) *Game {
	newGrid := initializeGrid(gridRows, gridCols)

	return &Game{
		server:    server,
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
				position: Position{
					x: uint8(colIndex),
					y: uint8(rowIndex),
				},
				content: Content{
					contentType: empty,
				},
			}
		}
	}
	return newGridState
}
