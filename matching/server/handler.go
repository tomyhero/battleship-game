package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/matching/data"
	"golang.org/x/net/websocket"
)

type Handler struct {
	server *Server
}

func (self Handler) Search(conn *websocket.Conn, data data.Interface) {

	fmt.Println(data, "Search", self.server.Waitings)

	enemyConn := &websocket.Conn{}
	onMatch := false
	self.server.WaitingLock.Lock()

	if len(self.server.Waitings) > 0 {
		enemyConn = self.server.Waitings[0]
		self.server.DeleteWaitingEntry(enemyConn)
		onMatch = true
	} else {
		self.server.Waitings = append(self.server.Waitings, conn)
	}

	self.server.WaitingLock.Unlock()

	if onMatch {
		// TODO send to matchid to enemy(to server)
		fmt.Println("EnemyConn", enemyConn)
		// return match (to client)
	}

}

func (seslf Handler) Found(conn *websocket.Conn, data data.Interface) {
	fmt.Println(data, "Found")

}
