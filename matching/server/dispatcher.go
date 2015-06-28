package server

import (
	"fmt"
	"github.com/tomyhero/battleship-game/matching/data"
	"github.com/tomyhero/battleship-game/matching/handler"
	"reflect"
)

type Dispatcher struct {
	actions map[string]reflect.Value
}

func NewDispatcher() Dispatcher {
	dispatcher := Dispatcher{}
	dispatcher.loadActions()
	return dispatcher
}

func (self *Dispatcher) loadActions() {
	actions := map[string]reflect.Value{}

	handler := handler.Handler{}
	actions["search"] = reflect.ValueOf(handler).MethodByName("Search")
	self.actions = actions
}

func (self *Dispatcher) findAction(cmd string) (reflect.Value, bool) {
	action, ok := self.actions[cmd]
	return action, ok
}

func (self *Dispatcher) Dispatch(d map[string]interface{}) error {
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
	action.Call([]reflect.Value{reflect.ValueOf(data)})
	return nil
}

// TODO make it more smarter
func (self *Dispatcher) findData(cmd string) (data.Interface, error) {
	if cmd == "search" {
		return data.Search{}, nil
	}

	return nil, fmt.Errorf("NOT_FOUND")
}
