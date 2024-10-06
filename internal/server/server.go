package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Player struct {
	id   string
	name string
}

const (
	player_limit = 5000
)

type GameMove struct {
	Action  string `json:"action"`
	Payload string `json:"payload"`
}

/**
* Primary struct on the websocket server instance and its performance.
**/
type Server struct {
	ListenAddr  string
	upgrader    websocket.Upgrader
	players     map[string]Player          // all players that can play
	clientConns map[*websocket.Conn]Player // all currently connected players from all match connections
	serverChan  chan GameMove
}

func NewServer(listenAddr string) *Server {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections by default for simplicity; you can add more logic here
			return true
		},
	}
	return &Server{
		ListenAddr: listenAddr,
		upgrader:   upgrader,
		// TODO: update this to persist from DB
		players: make(map[string]Player, player_limit),
	}
}
