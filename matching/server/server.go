package server

import (
	"golang.org/x/net/websocket"
	"net/http"
)

type Server struct {
	Conns map[string]*websocket.Conn
}

func (server *Server) ListenAndServe() {
	http.Handle("/matching", websocket.Handler(server.handler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (server *Server) handler(ws *websocket.Conn) {

}
