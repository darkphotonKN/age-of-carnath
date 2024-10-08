package server

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMultiplayerServer_FindMatch(t *testing.T) {
	testMultiplayerServer := NewMultiplayerServer(":3636")

	// setup mock matches
	mockMatches := make(map[uuid.UUID][]Player)

	// mockPlayerOneId := uuid.New()
	mockPlayerTwoId := uuid.New()
	testPlayerId := uuid.New()
	// gameOneId := uuid.New()
	gameTwoId := uuid.New()

	// creating one full game
	// mockMatches[gameOneId] = append(mockMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	// mockMatches[gameOneId] = append(mockMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 2"})
	// creating one half full game
	mockMatches[gameTwoId] = append(mockMatches[gameTwoId], Player{id: mockPlayerTwoId, name: "Mock Player 2"})

	// pre-feed matches with these two games
	testMultiplayerServer.matches = mockMatches

	// setup test player to be added
	testPlayer := Player{
		id:   testPlayerId,
		name: "TEST PLAYER",
	}

	// find a new match to test match find logic
	testMultiplayerServer.findMatch(testPlayer)

	expectedMatches := make(map[uuid.UUID][]Player)
	// expectedMatches[gameOneId] = append(expectedMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	// expectedMatches[gameOneId] = append(expectedMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 2"})
	expectedMatches[gameTwoId] = append(expectedMatches[gameTwoId], Player{id: mockPlayerTwoId, name: "Mock Player 2"})
	expectedMatches[gameTwoId] = append(expectedMatches[gameTwoId], testPlayer)

	fmt.Printf("EXPECTED\n\n")

	PrettyPrintMatches(expectedMatches)

	fmt.Printf("ACTUAL\n\n")

	PrettyPrintMatches(testMultiplayerServer.matches)

	assert.Equal(t, expectedMatches, testMultiplayerServer.matches)
}
