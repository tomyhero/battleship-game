package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/utils"
	"golang.org/x/net/websocket"
	"io"
	"net"
	"net/http"
)

type Server struct {
	Config *utils.Config
}

func NewServer() Server {
	server := Server{}
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

		fmt.Println("Dispatch", data)
		err = dispatcher.Dispatch(conn, data)

		if err != nil {
			fmt.Println(fmt.Sprintf("Dispatch Failed %s", err))
		}
	}

	defer func() {
		conn.Close()
	}()

}
