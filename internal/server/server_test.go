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

	expectedPrint := MapIdStringMatches(expectedMatches)

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

	// empty match slot since defaulting the match lengths to be 2, the real match will have a slot with an empty player too
	expectedMatches[matchFoundId] = append(expectedMatches[matchFoundId], Player{})

	expectedPrint := MapIdStringMatches(expectedMatches)

	actualPrint := MapIdStringMatches(testMultiplayerServer.matches)

	assert.Equal(t, expectedPrint, actualPrint)
}
