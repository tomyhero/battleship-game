package server

import (
	"fmt"
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
				fmt.Println(err)
				continue
			}
			if err == io.EOF {
				fmt.Println(fmt.Sprintf("Client Dissconected :%s", conn.RemoteAddr()))
				break
			} else {
				fmt.Println(fmt.Sprintf("Receive Data Failed %s", err))
				break
			}
		}

		err = dispatcher.Dispatch(conn, data)

		if err != nil {
			fmt.Println(fmt.Sprintf("Dispatch Failed %s", err))
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
		self.Waitings = append(self.Waitings[:match], self.Waitings[match+1:]...)
	}
}
