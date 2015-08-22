package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/tomyhero/battleship-game/matching/in"
	"github.com/tomyhero/battleship-game/utils"
	"golang.org/x/net/websocket"
)

type Handler struct {
	server *Server
}

func (self Handler) Search(conn *websocket.Conn, in in.Interface) {

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
		matchingID := utils.GUID()

		a := map[string]string{"cmd": "found", "user_id": utils.GUID(), "matching_id": matchingID}
		err := websocket.JSON.Send(enemyConn, a)

		if err != nil {
			log.WithFields(log.Fields{"error": err, "data": a}).Info("Fail To Send")
		}

		b := map[string]string{"cmd": "found", "user_id": utils.GUID(), "matching_id": matchingID}
		err = websocket.JSON.Send(conn, b)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "data": b}).Info("Fail To Send")
		}
	}

}
