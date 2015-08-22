package server

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/tomyhero/battleship-game/matching/in"
	"golang.org/x/net/websocket"
	"reflect"
)

type Dispatcher struct {
	actions map[string]reflect.Value
	server  *Server
}

func NewDispatcher(server *Server) Dispatcher {
	dispatcher := Dispatcher{server: server}
	dispatcher.loadActions()
	return dispatcher
}

func (self *Dispatcher) loadActions() {
	actions := map[string]reflect.Value{}

	handler := Handler{server: self.server}
	actions["search"] = reflect.ValueOf(handler).MethodByName("Search")
	actions["found"] = reflect.ValueOf(handler).MethodByName("Found")
	self.actions = actions
}

func (self *Dispatcher) findAction(cmd string) (reflect.Value, bool) {
	action, ok := self.actions[cmd]
	return action, ok
}

func (self *Dispatcher) Dispatch(conn *websocket.Conn, d map[string]interface{}) error {

	cmd, ok := d["cmd"]

	if !ok {
		return fmt.Errorf("command not found")
	}

	log.WithFields(log.Fields{"cmd": cmd}).Info("Dispatch To ")

	action, hasCommand := self.findAction(cmd.(string))

	if !hasCommand {
		return fmt.Errorf("action not found")
	}

	in := self.findIn(cmd.(string))
	in.Load(d)
	action.Call([]reflect.Value{reflect.ValueOf(conn), reflect.ValueOf(in)})
	return nil
}

func (self *Dispatcher) findIn(cmd string) in.Interface {
	if cmd == "search" {
		return &in.Search{}
	}
	log.WithFields(log.Fields{"cmd": cmd}).Fatal("Please Implement")
	return nil
}
