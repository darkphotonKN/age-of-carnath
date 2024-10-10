package server

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// -- Half Full Matches Test --
func TestNewMultiplayerServer_FindMatch_HalfFull(t *testing.T) {
	testMultiplayerServer := NewMultiplayerServer(":3636")

	// setup mock matches
	mockMatches := make(map[uuid.UUID][]Player)

	mockPlayerOneId := uuid.New()
	mockPlayerTwoId := uuid.New()
	mockPlayerThreeId := uuid.New()
	testPlayerId := uuid.New()
	gameOneId := uuid.New()
	gameTwoId := uuid.New()

	// setup test player to be added
	testPlayer := Player{
		id:   testPlayerId,
		name: "TEST PLAYER",
	}

	mockMatches[gameOneId] = append(mockMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	mockMatches[gameOneId] = append(mockMatches[gameOneId], Player{id: mockPlayerTwoId, name: "Mock Player 2"})
	mockMatches[gameTwoId] = append(mockMatches[gameTwoId], Player{id: mockPlayerThreeId, name: "Mock Player 3"})

	// pre-feed matches with these two games
	testMultiplayerServer.matches = mockMatches

	// simulate find a new match and test the results
	matchFoundId := testMultiplayerServer.findMatch(testPlayer)
	fmt.Println("matchFoundId:", matchFoundId)

	expectedMatches := make(map[uuid.UUID][]Player)
	expectedMatches[gameOneId] = append(expectedMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	expectedMatches[gameOneId] = append(expectedMatches[gameOneId], Player{id: mockPlayerTwoId, name: "Mock Player 2"})
	expectedMatches[gameTwoId] = append(expectedMatches[gameTwoId], Player{id: mockPlayerThreeId, name: "Mock Player 3"})
	expectedMatches[gameTwoId] = append(expectedMatches[gameTwoId], testPlayer)

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
	mockMatches := make(map[uuid.UUID][]Player)

	mockPlayerOneId := uuid.New()
	mockPlayerTwoId := uuid.New()
	mockPlayerThreeId := uuid.New()
	mockPlayerFourId := uuid.New()
	testPlayerId := uuid.New()
	gameOneId := uuid.New()
	gameTwoId := uuid.New()

	// setup test player to be added
	testPlayer := Player{
		id:   testPlayerId,
		name: "TEST PLAYER",
	}

	mockMatches[gameOneId] = append(mockMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	mockMatches[gameOneId] = append(mockMatches[gameOneId], Player{id: mockPlayerTwoId, name: "Mock Player 2"})
	mockMatches[gameTwoId] = append(mockMatches[gameTwoId], Player{id: mockPlayerThreeId, name: "Mock Player 3"})
	mockMatches[gameTwoId] = append(mockMatches[gameTwoId], Player{id: mockPlayerFourId, name: "Mock Player 4"})

	// pre-feed matches with these two games
	testMultiplayerServer.matches = mockMatches

	// simulate find a new match and test the results
	matchFoundId := testMultiplayerServer.findMatch(testPlayer)

	expectedMatches := make(map[uuid.UUID][]Player)

	fmt.Printf("matchFoundId: %v, existing slice: %+v\n", matchFoundId, expectedMatches[matchFoundId])
	fmt.Println("matchFoundId type:", reflect.TypeOf(matchFoundId))

	if expectedMatches[matchFoundId] == nil {
		expectedMatches[matchFoundId] = []Player{}
	}

	expectedMatches[gameOneId] = append(expectedMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	expectedMatches[gameOneId] = append(expectedMatches[gameOneId], Player{id: mockPlayerTwoId, name: "Mock Player 2"})
	expectedMatches[gameTwoId] = append(expectedMatches[gameTwoId], Player{id: mockPlayerThreeId, name: "Mock Player 3"})
	expectedMatches[gameTwoId] = append(expectedMatches[gameTwoId], Player{id: mockPlayerFourId, name: "Mock Player 4"})
	expectedMatches[matchFoundId] = append(expectedMatches[matchFoundId], testPlayer)

	expectedPrint := MapIdStringMatches(expectedMatches)

	actualPrint := MapIdStringMatches(testMultiplayerServer.matches)

	assert.Equal(t, expectedPrint, actualPrint)
}

// -- Multiple player joins via findMatch Test --
func TestNewMultiplayerServer_FindMatch_Multiple(t *testing.T) {
	fmt.Println("Testing Find Match --- Multiple Matches")
	testMultiplayerServer := NewMultiplayerServer(":3636")

	// setup mock matches
	mockMatches := make(map[uuid.UUID][]Player)

	mockPlayerOneId := uuid.New()
	testPlayerId := uuid.New()
	testPlayerTwoId := uuid.New()
	testPlayerThreeId := uuid.New()
	gameOneId := uuid.New()

	// setup test players to be added
	testPlayer := Player{
		id:   testPlayerId,
		name: "TEST PLAYER",
	}

	testPlayer2 := Player{
		id:   testPlayerTwoId,
		name: "TEST PLAYER 2",
	}

	testPlayer3 := Player{
		id:   testPlayerThreeId,
		name: "TEST PLAYER 3",
	}

	mockMatches[gameOneId] = append(mockMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})

	// pre-feed matches with these two games
	testMultiplayerServer.matches = mockMatches

	// simulate MULTIPLE findMatch and test the results
	matchFoundIdOne := testMultiplayerServer.findMatch(testPlayer)
	matchFoundIdTwo := testMultiplayerServer.findMatch(testPlayer2)
	testMultiplayerServer.findMatch(testPlayer3)

	expectedMatches := make(map[uuid.UUID][]Player)
	expectedMatches[gameOneId] = append(expectedMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	expectedMatches[matchFoundIdOne] = append(expectedMatches[matchFoundIdOne], testPlayer)
	expectedMatches[matchFoundIdTwo] = append(expectedMatches[matchFoundIdTwo], testPlayer2)
	expectedMatches[matchFoundIdTwo] = append(expectedMatches[matchFoundIdTwo], testPlayer3)

	fmt.Print("EXPECTED: ")
	expectedPrint := MapIdStringMatches(expectedMatches)

	fmt.Print("ACTUAL: ")
	actualPrint := MapIdStringMatches(testMultiplayerServer.matches)

	assert.Equal(t, expectedPrint, actualPrint)
}
