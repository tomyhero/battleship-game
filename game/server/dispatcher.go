package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/game/in"
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
	actions["start"] = reflect.ValueOf(handler).MethodByName("Start")
	actions["attack"] = reflect.ValueOf(handler).MethodByName("Attack")
	actions["login"] = reflect.ValueOf(handler).MethodByName("Login")
	actions["resume"] = reflect.ValueOf(handler).MethodByName("Resume")

	self.actions = actions
}

// TODO make it more smarter
func (self *Dispatcher) findIn(cmd string) (in.Interface, error) {
	if cmd == "start" {
		return &in.Start{}, nil
	} else if cmd == "attack" {
		return &in.Attack{}, nil
	} else if cmd == "login" {
		return &in.Login{}, nil
	} else if cmd == "resume" {
		return &in.Resume{}, nil
	}
	return nil, fmt.Errorf("NOT_FOUND")
}

func (self *Dispatcher) findAction(cmd string) (reflect.Value, bool) {
	action, ok := self.actions[cmd]
	return action, ok
}

func (self *Dispatcher) Dispatch(conn *websocket.Conn, d map[string]interface{}) error {

	cmd, ok := d["cmd"]

	if !ok {
		fmt.Println("does not have cmd section")
		return fmt.Errorf("CMD_NOT_FOUND")
	}
	action, hasCommand := self.findAction(cmd.(string))

	if !hasCommand {
		return fmt.Errorf("NOT_FOUND")
	}

	in, _ := self.findIn(cmd.(string))

	in.Load(d)

	action.Call([]reflect.Value{reflect.ValueOf(conn), reflect.ValueOf(in)})
	return nil
}
