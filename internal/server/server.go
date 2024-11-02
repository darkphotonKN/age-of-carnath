package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/darkphotonKN/age-of-carnath/internal/game"
	"github.com/darkphotonKN/age-of-carnath/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	player_limit = 5000
)

// to group GameMessage along with the websocket connection to pass to message hub for handling
type ClientPackage struct {
	GameMessage GameMessage
	Conn        *websocket.Conn
}

/**
* Primary struct on the websocket server instance and its performance.
**/
type MultiplayerServer struct {
	ListenAddr string
	upgrader   websocket.Upgrader

	// All players. TODO: Update this to use a persistent database.
	players map[string]models.Player

	// All currently connected players from all on-going match connections
	clientConns map[*websocket.Conn]models.Player

	// Communication for each game - MAP of channels of GameMessage
	gameMessageChans map[*websocket.Conn]chan GameMessage

	// All ongoing matches
	// NOTE: self-reminder - maintain this as a map for O(1) performance in finding matches
	matches    map[uuid.UUID]*game.Game
	serverChan chan ClientPackage
	mu         sync.Mutex
}

type Math struct {
	value int
}

func (m Math) Sum(x1, x2 int) int {
	m.value = x1 + x2
	return m.value
}

func (m Math) Subtract(x1, x2 int) int {
	m.value = x1 - x2
	return m.value
}

func NewMultiplayerServer(listenAddr string) *MultiplayerServer {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// TODO: Allow all connections by default for simplicity; can add more logic here
			return true
		},
	}

	mathHelper := Math{}
	mathHelper.Sum(1, 2)

	return &MultiplayerServer{
		ListenAddr:       listenAddr,
		upgrader:         upgrader,
		players:          make(map[string]models.Player, player_limit),
		clientConns:      make(map[*websocket.Conn]models.Player, player_limit),
		gameMessageChans: make(map[*websocket.Conn]chan GameMessage),
		matches:          make(map[uuid.UUID]*game.Game),
		serverChan:       make(chan ClientPackage),
	}
}

// --- Helpers ---

/**
* Thread-safe lock-unlock access of the game message channels used for writing
* back to clients.
*
**/
func (s *MultiplayerServer) getGameMsgChan(conn *websocket.Conn) chan GameMessage {
	s.mu.Lock()
	defer s.mu.Unlock()
	channel, exists := s.gameMessageChans[conn]

	if exists {
		return channel
	}

	return nil
}

/**
* Creates a game message channel if it doesn't exist
**/
func (s *MultiplayerServer) createGameMsgChan(conn *websocket.Conn) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.gameMessageChans[conn]

	if exists {
		return fmt.Errorf("Game message channel already exist.")
	}

	// create channel and add it to the map
	newChan := make(chan GameMessage)
	s.gameMessageChans[conn] = newChan

	return nil
}
