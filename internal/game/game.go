package game

import "github.com/darkphotonKN/age-of-carnath/internal/server"

/**
* Position on the Grid
**/
type Position struct {
	x uint8
	y uint8
}

/**
* Represents a Single Grid Block
* Position - the occupying position of this block in relation to the entire grid.
* Content - the occupying content of the grid.
**/
type GridBlock struct {
	position Position
	content  interface{}
}

/**
* Represents the entire game grid, compopsed of GridBlocks.
**/
type GridState [][]GridBlock

type Game struct {
	server    *server.Server
	gridState GridState
}

func NewGame() *Game {

	return &Game{}
}
