package server

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/tomyhero/battleship-game/utils"
	"golang.org/x/net/websocket"
	"io"
	"net"
	"net/http"
	"sync"
)

type Server struct {
	Config      *utils.Config
	Waitings    []*websocket.Conn
	WaitingLock sync.RWMutex
}

func NewServer() Server {
	server := Server{}
	return server
}

func (self *Server) ListenAndServe(port int) {
	self.Waitings = []*websocket.Conn{}
	http.Handle("/matching", websocket.Handler(self.webSocketHandler))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Fatal("Listen And Serve")
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
			log.WithFields(log.Fields{"err": err}).Warn("Dispatch Failed")
		}
	}

	defer func() {
		self.WaitingLock.Lock()

		self.DeleteWaitingEntry(conn)
		conn.Close()

		self.WaitingLock.Unlock()
	}()

}

func (self *Server) DeleteWaitingEntry(conn *websocket.Conn) {
	match := -1
	for i := 0; i < len(self.Waitings); i++ {
		if self.Waitings[i] == conn {
			match = i
		}
	}
	if match != -1 {
		log.Info("Delete Waiting Entry")
		self.Waitings = append(self.Waitings[:match], self.Waitings[match+1:]...)
	}
}
