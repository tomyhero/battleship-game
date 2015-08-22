package server

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/tomyhero/battleship-game/game/model"
	"github.com/tomyhero/battleship-game/utils"
	"golang.org/x/net/websocket"
	"io"
	"net"
	"net/http"
	"sync"
)

type Server struct {
	Config         *utils.Config
	Rooms          map[string]*model.Room
	Users          map[*websocket.Conn]*model.User
	disconnectLock sync.RWMutex
}

func NewServer() Server {
	server := Server{Rooms: map[string]*model.Room{}, Users: map[*websocket.Conn]*model.User{}}
	return server
}

func (self *Server) ListenAndServe(port int) {
	http.Handle("/game", websocket.Handler(self.webSocketHandler))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (self *Server) webSocketHandler(conn *websocket.Conn) {

	dispatcher := NewDispatcher(self)

	for {
		data := map[string]interface{}{}
		err := websocket.JSON.Receive(conn, &data)

		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				log.WithFields(log.Fields{"error": err}).Info("Error")
				continue
			}
			if err == io.EOF {
				log.WithFields(log.Fields{"remoteAddr": conn.RemoteAddr()}).Info("Client Dissconected")
				break
			} else {
				log.WithFields(log.Fields{"error": err}).Info("Receive Data Failed")
				break
			}
		}

		log.WithFields(log.Fields{"data": data}).Info("Start Dispatch")
		err = dispatcher.Dispatch(conn, data)

		if err != nil {
			fmt.Println(fmt.Sprintf("Dispatch Failed %s", err))
		}
	}

	defer func() {
		self.disconnect(conn)
		conn.Close()
	}()

}

func (self *Server) disconnect(conn *websocket.Conn) {
	self.disconnectLock.Lock()
	user, has := self.Users[conn]
	if has {
		log.WithFields(log.Fields{"MatchingID": user.MatchingID, "UserID": user.UserID}).Info("Delete User From server.Users Entry")

		delete(self.Users, conn)
		room, has := self.Rooms[user.MatchingID]
		if has {
			user.IsConnected = false
			user.Conn = nil
			if room.Users[room.EnemyUserID(user.UserID)].IsConnected == false {
				log.WithFields(log.Fields{"MatchingID": room.MatchingID}).Info("Delete Room From server.Rooms Entry")
				delete(self.Rooms, user.MatchingID)
			}
		}
	}
	self.disconnectLock.Unlock()
}
