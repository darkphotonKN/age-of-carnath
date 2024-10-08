package server

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMultiplayerServer_FindMatch(t *testing.T) {
	testMultiplayerServer := NewMultiplayerServer(":3636")

	// setup mock matches
	mockMatches := make(map[uuid.UUID][]Player)

	mockPlayerOneId := uuid.New()
	mockPlayerTwoId := uuid.New()
	testPlayerId := uuid.New()
	gameOneId := uuid.New()
	gameTwoId := uuid.New()

	// creating two half empty games
	mockMatches[gameOneId] = append(mockMatches[gameOneId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	mockMatches[gameTwoId] = append(mockMatches[gameTwoId], Player{id: mockPlayerTwoId, name: "Mock Player 2"})

	testMultiplayerServer.matches = mockMatches

	// setup test player to be added
	testPlayer := Player{
		id:   testPlayerId,
		name: "TEST PLAYER",
	}

	// find a new match to test match find logic
	testMultiplayerServer.findMatch(testPlayer)

	expectedMatches := make(map[uuid.UUID][]Player)
	expectedMatches[gameOneId] = append(expectedMatches[gameTwoId], Player{id: mockPlayerOneId, name: "Mock Player 1"})
	expectedMatches[gameTwoId] = append(expectedMatches[gameTwoId], Player{id: mockPlayerTwoId, name: "Mock Player 2"})
	expectedMatches[uuid.New()] = append(expectedMatches[uuid.New()], Player{id: testPlayerId, name: "TEST PLAYER"})

	assert.Equal(t, expectedMatches, testMultiplayerServer.matches)

}
