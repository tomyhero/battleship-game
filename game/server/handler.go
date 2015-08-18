package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/game/data"
	"golang.org/x/net/websocket"
	"sync"
)

type Handler struct {
	server    *Server
	StartLock sync.RWMutex
}

func (self Handler) Start(conn *websocket.Conn, d data.Interface) {
	in := d.(*data.Start)
	fmt.Println(in)
	self.StartLock.Lock()

	room, has := self.server.Rooms[in.MatchingID]

	if has {
		fmt.Println("room", room)
	} else {
		room = data.NewRoom()
		room.Users[in.UserID] = data.User{}
		self.server.Rooms[in.MatchingID] = room
	}

	self.StartLock.Unlock()
}
