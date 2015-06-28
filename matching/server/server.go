package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/utils"
	"golang.org/x/net/websocket"
	"net/http"
)

type Server struct {
	Config *utils.Config
	Conns  map[*websocket.Conn]int
}

func (self *Server) ListenAndServe(port int) {
	self.Conns = map[*websocket.Conn]int{}
	http.Handle("/matching", websocket.Handler(self.handler))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (self *Server) handler(conn *websocket.Conn) {

	self.Conns[conn] = 1

	defer func() {
		conn.Close()
		fmt.Println(self.Conns)
		delete(self.Conns, conn)
		fmt.Println(self.Conns)
		fmt.Println("Close")
	}()

}
