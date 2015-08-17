package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/matching/data"
	"golang.org/x/net/websocket"
)

type Handler struct {
	server *Server
}

func (self Handler) Start(conn *websocket.Conn, data data.Interface) {
	fmt.Println(data, "Start")
}
