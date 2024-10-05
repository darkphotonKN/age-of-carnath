package server

import "github.com/gorilla/websocket"

type Player struct {
	id   string
	name string
}

const (
	player_limit = 5000
)

/**
* Primarily on the websocket server instance and its performance.
**/
type Server struct {
	listenAddr  string
	players     map[string]Player          // all players that can play
	clientConns map[*websocket.Conn]Player // all currently connected players from all match connections
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		// TODO: update this to persist from DB
		players: make(map[string]Player, player_limit),
	}
}
