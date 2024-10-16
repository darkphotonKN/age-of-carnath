package server

import (
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

	// All ongoing matches
	// NOTE: self-reminder - maintain this as a map for O(1) performance in finding matches
	matches    map[uuid.UUID]*game.Game
	serverChan chan ClientPackage
	mu         sync.Mutex
}

func NewMultiplayerServer(listenAddr string) *MultiplayerServer {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// TODO: Allow all connections by default for simplicity; can add more logic here
			return true
		},
	}

	return &MultiplayerServer{
		ListenAddr:  listenAddr,
		upgrader:    upgrader,
		players:     make(map[string]models.Player, player_limit),
		clientConns: make(map[*websocket.Conn]models.Player, player_limit),
		matches:     make(map[uuid.UUID]*game.Game),
		serverChan:  make(chan ClientPackage),
	}
}
