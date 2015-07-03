package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/matching/data"
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

	fmt.Println("Call Dispatch", d)
	cmd, ok := d["cmd"]

	if !ok {
		fmt.Println("does not have cmd section")
		return fmt.Errorf("CMD_NOT_FOUND")
	}
	action, hasCommand := self.findAction(cmd.(string))

	if !hasCommand {
		return fmt.Errorf("NOT_FOUND")
	}

	data, _ := self.findData(cmd.(string))
	data.Load(d)
	action.Call([]reflect.Value{reflect.ValueOf(conn), reflect.ValueOf(data)})
	return nil
}

// TODO make it more smarter
func (self *Dispatcher) findData(cmd string) (data.Interface, error) {
	if cmd == "search" {
		return data.Search{}, nil
	} else if cmd == "found" {
		return data.Found{}, nil
	}

	return nil, fmt.Errorf("NOT_FOUND")
}
