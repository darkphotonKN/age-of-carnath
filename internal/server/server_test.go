package server

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/darkphotonKN/age-of-carnath/internal/game"
	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// -- Half Full Matches Test --
func TestNewMultiplayerServer_FindMatch_HalfFull(t *testing.T) {
	testMultiplayerServer := NewMultiplayerServer(":3636")

	// setup mock matches
	mockMatches := make(map[uuid.UUID]*game.Game)

	mockPlayerOneId := uuid.New()
	mockPlayerTwoId := uuid.New()
	mockPlayerThreeId := uuid.New()
	testPlayerId := uuid.New()
	gameOneId := uuid.New()
	gameTwoId := uuid.New()

	// setup test player to be added
	testPlayer := models.Player{
		ID:   testPlayerId,
		Name: "TEST PLAYER",
	}

	// initialize each game struct
	mockPlayer1 := models.Player{ID: mockPlayerOneId, Name: "Mock Player 1"}
	mockMatches[gameOneId] = game.InitializeGame(&mockPlayer1)

	mockPlayer2 := models.Player{ID: mockPlayerThreeId, Name: "Mock Player 3"}
	mockMatches[gameTwoId] = game.InitializeGame(&mockPlayer2)

	mockMatches[gameOneId].Players = append(mockMatches[gameOneId].Players, models.Player{ID: mockPlayerTwoId, Name: "Mock Player 2"})

	// pre-feed matches with these two games
	testMultiplayerServer.matches = mockMatches

	// simulate find a new match and test the results
	matchFoundId := testMultiplayerServer.findMatch(testPlayer)
	fmt.Println("matchFoundId:", matchFoundId)

	expectedMatches := make(map[uuid.UUID]*game.Game)

	// initialize each game struct
	expectedMatches[gameOneId] = &game.Game{}
	expectedMatches[gameTwoId] = &game.Game{}

	expectedMatches[gameOneId].Players = append(expectedMatches[gameOneId].Players, models.Player{ID: mockPlayerOneId, Name: "Mock Player 1"})
	expectedMatches[gameOneId].Players = append(expectedMatches[gameOneId].Players, models.Player{ID: mockPlayerTwoId, Name: "Mock Player 2"})
	expectedMatches[gameTwoId].Players = append(expectedMatches[gameTwoId].Players, models.Player{ID: mockPlayerThreeId, Name: "Mock Player 3"})
	expectedMatches[gameTwoId].Players = append(expectedMatches[gameTwoId].Players, testPlayer)

	fmt.Print("EXPECTED: ")
	expectedPrint := MapIdStringMatches(expectedMatches)

	fmt.Print("ACTUAL: ")
	actualPrint := MapIdStringMatches(testMultiplayerServer.matches)

	assert.Equal(t, expectedPrint, actualPrint)
}

// -- Full Matches Test
func TestNewMultiplayerServer_FindMatch_Full(t *testing.T) {
	fmt.Println("Testing Find Match --- Full Matches")
	testMultiplayerServer := NewMultiplayerServer(":3636")

	// setup mock matches
	mockMatches := make(map[uuid.UUID]*game.Game)

	mockPlayerOneId := uuid.New()
	mockPlayerTwoId := uuid.New()
	mockPlayerThreeId := uuid.New()
	mockPlayerFourId := uuid.New()
	testPlayerId := uuid.New()
	gameOneId := uuid.New()
	gameTwoId := uuid.New()

	// setup test player to be added
	testPlayer := models.Player{
		ID:   testPlayerId,
		Name: "TEST PLAYER",
	}

	// initialize each game struct
	mockPlayer1 := models.Player{ID: mockPlayerOneId, Name: "Mock Player 1"}
	mockMatches[gameOneId] = game.InitializeGame(&mockPlayer1)
	mockPlayer3 := models.Player{ID: mockPlayerThreeId, Name: "Mock Player 3"}
	mockMatches[gameTwoId] = game.InitializeGame(&mockPlayer3)

	mockMatches[gameOneId].Players = append(mockMatches[gameOneId].Players, models.Player{ID: mockPlayerTwoId, Name: "Mock Player 2"})
	mockMatches[gameTwoId].Players = append(mockMatches[gameTwoId].Players, models.Player{ID: mockPlayerFourId, Name: "Mock Player 4"})

	// pre-feed matches with these two games
	testMultiplayerServer.matches = mockMatches

	// simulate find a new match and test the results
	matchFoundId := testMultiplayerServer.findMatch(testPlayer)

	expectedMatches := make(map[uuid.UUID]*game.Game)

	// initialize each game struct
	expectedMatches[gameOneId] = &game.Game{}
	expectedMatches[gameTwoId] = &game.Game{}
	expectedMatches[matchFoundId] = &game.Game{}

	fmt.Printf("matchFoundId: %v, existing slice: %+v\n", matchFoundId, expectedMatches[matchFoundId])
	fmt.Println("matchFoundId type:", reflect.TypeOf(matchFoundId))

	expectedMatches[gameOneId].Players = append(expectedMatches[gameOneId].Players, models.Player{ID: mockPlayerOneId, Name: "Mock Player 1"})
	expectedMatches[gameOneId].Players = append(expectedMatches[gameOneId].Players, models.Player{ID: mockPlayerTwoId, Name: "Mock Player 2"})
	expectedMatches[gameTwoId].Players = append(expectedMatches[gameTwoId].Players, models.Player{ID: mockPlayerThreeId, Name: "Mock Player 3"})
	expectedMatches[gameTwoId].Players = append(expectedMatches[gameTwoId].Players, models.Player{ID: mockPlayerFourId, Name: "Mock Player 4"})
	expectedMatches[matchFoundId].Players = append(expectedMatches[matchFoundId].Players, testPlayer)

	expectedPrint := MapIdStringMatches(expectedMatches)

	actualPrint := MapIdStringMatches(testMultiplayerServer.matches)

	assert.Equal(t, expectedPrint, actualPrint)
}

// -- Multiple player joins via findMatch Test --
func TestNewMultiplayerServer_FindMatch_Multiple(t *testing.T) {
	fmt.Println("Testing Find Match --- Multiple Matches")
	testMultiplayerServer := NewMultiplayerServer(":3636")

	// setup mock matches
	mockMatches := make(map[uuid.UUID]*game.Game)

	mockPlayerOneId := uuid.New()
	testPlayerId := uuid.New()
	testPlayerTwoId := uuid.New()
	testPlayerThreeId := uuid.New()
	gameOneId := uuid.New()

	// initialize each game struct
	mockPlayer1 := models.Player{ID: mockPlayerOneId, Name: "Mock Player 1"}
	mockMatches[gameOneId] = game.InitializeGame(&mockPlayer1)

	// setup test players to be added
	testPlayer := models.Player{
		ID:   testPlayerId,
		Name: "TEST PLAYER",
	}

	testPlayer2 := models.Player{
		ID:   testPlayerTwoId,
		Name: "TEST PLAYER 2",
	}

	testPlayer3 := models.Player{
		ID:   testPlayerThreeId,
		Name: "TEST PLAYER 3",
	}

	// pre-feed matches with these two games
	testMultiplayerServer.matches = mockMatches

	// simulate MULTIPLE findMatch and test the results
	matchFoundIdOne := testMultiplayerServer.findMatch(testPlayer)
	matchFoundIdTwo := testMultiplayerServer.findMatch(testPlayer2)
	testMultiplayerServer.findMatch(testPlayer3)

	expectedMatches := make(map[uuid.UUID]*game.Game)
	expectedMatches[gameOneId] = &game.Game{}
	expectedMatches[matchFoundIdOne] = &game.Game{}
	expectedMatches[matchFoundIdTwo] = &game.Game{}

	expectedMatches[gameOneId].Players = append(expectedMatches[gameOneId].Players, models.Player{ID: mockPlayerOneId, Name: "Mock Player 1"})
	expectedMatches[matchFoundIdOne].Players = append(expectedMatches[matchFoundIdOne].Players, testPlayer)
	expectedMatches[matchFoundIdTwo].Players = append(expectedMatches[matchFoundIdTwo].Players, testPlayer2)
	expectedMatches[matchFoundIdTwo].Players = append(expectedMatches[matchFoundIdTwo].Players, testPlayer3)

	fmt.Print("EXPECTED: ")
	expectedPrint := MapIdStringMatches(expectedMatches)

	fmt.Print("ACTUAL: ")
	actualPrint := MapIdStringMatches(testMultiplayerServer.matches)

	assert.Equal(t, expectedPrint, actualPrint)
}
