package server

import (
	"fmt"
	"sync"

	"github.com/tomyhero/battleship-game/game/in"
	"github.com/tomyhero/battleship-game/game/model"
	"golang.org/x/net/websocket"
)

type Handler struct {
	server    *Server
	StartLock sync.RWMutex
}

func (self Handler) Attack(conn *websocket.Conn, d in.Interface) {
	in := d.(*in.Attack)
	room, has := self.server.Rooms[in.MatchingID]
	if !has {
		fmt.Println("Room Not Found")
		return
	}

	userID, _, err := room.GetUserFromConn(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !room.IsYourTurn(userID) {
		fmt.Println("Not Your Turn")
		return
	}

	enemy, has := room.Users[room.EnemyUserID(userID)]
	if !has {
		fmt.Println("Enemy Not Found")
		return
	}

	field := enemy.Fields[in.Y][in.X]

	if field.HitType != model.HIT_TYPE.YET {
		fmt.Println("Already Attacked")
		return
	}

	if field.ShipID != 0 {
		field.HitType = model.HIT_TYPE.MISS
	} else if enemy.IsNearShip(in.Y, in.X) {
		field.HitType = model.HIT_TYPE.NEAR
	} else {
		field.HitType = model.HIT_TYPE.HIT
	}

	destroyShipSize := room.GetDestroyShipSize(enemy, field)

	fmt.Println("destoroy", destroyShipSize)

	room.ChangeTurn()

	onFinish := room.IsFinishGame(userID)

	data := map[string]interface{}{"x": in.X, "y": in.Y, "field": field, "on_finish": onFinish, "destroy_ship_size": destroyShipSize}

	err = websocket.JSON.Send(conn, map[string]interface{}{"cmd": "turn_end", "data": data})
	if err != nil {
		fmt.Println(err)
	}

	if onFinish {
		// lose
		data["enemyFields"] = room.Users[userID].Fields
	}

	err = websocket.JSON.Send(enemy.Conn, map[string]interface{}{"cmd": "turn_start", "data": data})
	if err != nil {
		fmt.Println(err)
	}

	if onFinish {
		delete(self.server.Rooms, in.MatchingID)
	}
}

func (self Handler) Start(conn *websocket.Conn, d in.Interface) {
	in := d.(*in.Start)
	self.StartLock.Lock()

	room, has := self.server.Rooms[in.MatchingID]

	if has {
		room.SetUser(in.UserID, conn)

		for userID, user := range room.Users {
			json := room.ToData(userID)
			err := websocket.JSON.Send(user.Conn, map[string]interface{}{"cmd": "start", "data": json})
			if err != nil {
				fmt.Println(err)
			}
		}

	} else {
		room = model.NewRoom()
		room.SetUser(in.UserID, conn)
		self.server.Rooms[in.MatchingID] = room
	}

	self.StartLock.Unlock()
}
