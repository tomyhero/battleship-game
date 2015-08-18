package model

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"math/rand"
	"sort"
)

type Room struct {
	Status      int
	Order       []string
	CurrentTurn int // index of Current Turn User
	Users       map[string]*User
	GameSetting GameSetting
}

var STATUS = NewCONST_STATUS()

type CONST_STATUS struct {
	INITIALIZE int
	IN_GAME    int
	FINISH     int
}

func NewCONST_STATUS() CONST_STATUS {
	c := CONST_STATUS{}
	c.INITIALIZE = 1
	c.IN_GAME = 2
	c.FINISH = 3

	return c
}

func NewRoom() *Room {
	room := Room{
		Users:       map[string]*User{},
		GameSetting: NewGameSetting(),
		Status:      STATUS.INITIALIZE,
		Order:       []string{}}
	return &room
}

func (self *Room) EnemyUserID(userID string) string {
	enemyUserID := ""
	for id, _ := range self.Users {
		if id != userID {
			enemyUserID = id
		}
	}
	return enemyUserID
}

func (self *Room) Enemy(userID string) map[string]interface{} {
	user := self.Users[self.EnemyUserID(userID)]

	fields := make([][]interface{}, self.GameSetting.MaxY)

	for y, line := range user.Fields {
		fields[y] = make([]interface{}, self.GameSetting.MaxX)
		for x, field := range line {
			fields[y][x] = map[string]interface{}{"HitType": field.HitType}
		}
	}
	return map[string]interface{}{"Fields": fields}
}

func (self *Room) ToJSON(userID string) string {
	d := map[string]interface{}{}
	d["Status"] = self.Status
	d["CurrentTurn"] = self.CurrentTurn
	d["Me"] = self.Users[userID]
	d["Enemy"] = self.Enemy(userID)

	json, _ := json.Marshal(d)

	return string(json)
}

type GameSetting struct {
	MaxX      int
	MaxY      int
	MaxPlayer int
	Ships     map[int]int
}

func NewGameSetting() GameSetting {
	g := GameSetting{MaxX: 16, MaxY: 16, MaxPlayer: 2}
	g.Ships = map[int]int{2: 4, 3: 3, 4: 2, 5: 1}
	return g
}

type User struct {
	Fields [][]*Field
	Conn   *websocket.Conn
}

type Field struct {
	HitType       int
	ShipID        int
	ShipPart      int
	ShipDirection bool
}

func (self *User) HideShips(g GameSetting) {

	shipSizes := []int{}

	for size, _ := range g.Ships {
		shipSizes = append(shipSizes, size)
	}

	// Hide from Big ship
	sort.Sort(sort.Reverse(sort.IntSlice(shipSizes)))
	id := 1
	for _, size := range shipSizes {
		count := g.Ships[size]
		for i := 0; i < count; i++ {
			HideShip(id, self.Fields, size)
			id = id + 1
		}
	}

}

func HideShip(id int, fields [][]*Field, size int) {

	// 起点と向きをランダムで求めて、配備できるまで繰り返す

	maxY := len(fields)
	maxX := len(fields[0])

	ok := false
	for !ok {
		direction := false

		if rand.Intn(2) == 1 {
			direction = true
		}

		// 横
		if direction == true {
			y := rand.Intn(maxY)
			x := rand.Intn(maxX - size)

			check := true
			for i := 0; i < size; i++ {
				field := fields[y][x+i]
				if field.ShipID != 0 {
					check = false
				}
			}

			if check {
				ok = true

				for i := 0; i < size; i++ {
					field := fields[y][x+i]

					field.ShipID = id
					field.ShipDirection = direction

					if i == 0 {
						field.ShipPart = 0
					} else if i == (size - 1) {
						field.ShipPart = 2
					} else {
						field.ShipPart = 1
					}

				}
			}

			// 縦
		} else {
			y := rand.Intn(maxY - size)
			x := rand.Intn(maxX)

			check := true
			for i := 0; i < size; i++ {
				field := fields[y+i][x]
				if field.ShipID != 0 {
					check = false
				}
			}

			if check {
				ok = true

				for i := 0; i < size; i++ {
					field := fields[y+i][x]

					field.ShipID = id
					field.ShipDirection = direction

					if i == 0 {
						field.ShipPart = 0
					} else if i == (size - 1) {
						field.ShipPart = 2
					} else {
						field.ShipPart = 1
					}

				}
			}

		}

	}
}

func NewUser(g GameSetting, conn *websocket.Conn) *User {
	user := &User{}
	user.Fields = g.GenerateFields()
	user.HideShips(g)
	user.Conn = conn
	return user
}

func (self *Room) SetUser(userID string, conn *websocket.Conn) {
	self.Users[userID] = NewUser(self.GameSetting, conn)
	self.Order = append(self.Order, userID)

}

func (self GameSetting) GenerateFields() [][]*Field {
	var fields [][]*Field
	fields = make([][]*Field, self.MaxY)

	for y := 0; y < self.MaxY; y++ {
		fields[y] = make([]*Field, self.MaxX)
		for x := 0; x < self.MaxX; x++ {
			fields[y][x] = &Field{}
		}
	}
	return fields
}
