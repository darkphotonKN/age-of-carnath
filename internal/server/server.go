package server

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
	listenAddr string
	players    map[string]Player
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		players:    make(map[string]Player, player_limit),
	}

}
