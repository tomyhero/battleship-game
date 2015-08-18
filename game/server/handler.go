package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/game/in"
	"github.com/tomyhero/battleship-game/game/model"
	"golang.org/x/net/websocket"
	"sync"
)

type Handler struct {
	server    *Server
	StartLock sync.RWMutex
}

func (self Handler) Start(conn *websocket.Conn, d in.Interface) {
	in := d.(*in.Start)
	self.StartLock.Lock()

	room, has := self.server.Rooms[in.MatchingID]

	if has {
		room.SetUser(in.UserID, conn)

		for userID, user := range room.Users {
			json := room.ToJSON(userID)
			err := websocket.JSON.Send(user.Conn, map[string]string{"cmd": "start", "data": json})
			if err != nil {
				fmt.Println(err)
			}
		}

	} else {
		room = model.NewRoom()
		room.SetUser(in.UserID, conn)
		self.server.Rooms[in.MatchingID] = room
	}

	self.StartLock.Unlock()
}
