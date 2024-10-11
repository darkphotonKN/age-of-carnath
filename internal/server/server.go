package server

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

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
	// TODO: Update this to database struct
	players     map[string]Player          // all players that can play
	clientConns map[*websocket.Conn]Player // all currently connected players from all match connections
	matches     map[uuid.UUID][]Player     // all ongoing matches
	serverChan  chan ClientPackage
	mu          sync.Mutex
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
		players:     make(map[string]Player, player_limit), // TODO: update this to persist from DB
		clientConns: make(map[*websocket.Conn]Player, player_limit),
		matches:     make(map[uuid.UUID][]Player),
		serverChan:  make(chan ClientPackage),
	}
}
