package model

import (
	//"encoding/json"
	"fmt"
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
var HIT_TYPE = NewCONST_HIT_TYPE()

type CONST_HIT_TYPE struct {
	YET  int
	MISS int
	NEAR int
	HIT  int
}

func NewCONST_HIT_TYPE() CONST_HIT_TYPE {
	c := CONST_HIT_TYPE{}
	c.YET = 0
	c.MISS = 1
	c.NEAR = 2
	c.HIT = 3
	return c
}

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

func (self *Room) IsFinishGame(userID string) bool {
	enemy := self.Users[self.EnemyUserID(userID)]

	for _, line := range enemy.Fields {
		for _, field := range line {
			if field.ShipID != 0 {
				// 壊れてない船がまだある
				if field.HitType == HIT_TYPE.YET {
					return false
				}
			}
		}
	}

	return true
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

func (self *Room) ChangeTurn() {

	if self.CurrentTurn == 0 {
		self.CurrentTurn = 1
	} else {
		self.CurrentTurn = 0
	}

}

func (self *Room) IsYourTurn(userID string) bool {
	id := self.Order[self.CurrentTurn]

	if userID == id {
		return true
	} else {
		return false
	}
}

func (self *Room) ToData(userID string) map[string]interface{} {
	d := map[string]interface{}{}
	d["Status"] = self.Status
	d["IsYourTurn"] = self.IsYourTurn(userID)
	d["Me"] = self.Users[userID]
	d["Enemy"] = self.Enemy(userID)
	d["Ships"] = self.GameSetting.ShipData()
	return d
}

type GameSetting struct {
	MaxX      int
	MaxY      int
	MaxPlayer int
	Ships     map[int]int
}

func (self GameSetting) ShipData() []map[string]interface{} {
	a := []map[string]interface{}{}

	shipSizes := []int{}

	for size, _ := range self.Ships {
		shipSizes = append(shipSizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(shipSizes)))

	for _, size := range shipSizes {
		count := self.Ships[size]
		d := map[string]interface{}{"size": size, "count": count}
		a = append(a, d)
	}
	return a
}

func NewGameSetting() GameSetting {
	g := GameSetting{MaxX: 16, MaxY: 16, MaxPlayer: 2}
	g.Ships = map[int]int{2: 4, 3: 3, 4: 2, 5: 1}
	//g := GameSetting{MaxX: 8, MaxY: 8, MaxPlayer: 2}
	//g.Ships = map[int]int{2: 1}
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

func (self *User) IsNearShip(y int, x int) bool {
	a := y + 1
	b := y - 1

	if a < len(self.Fields) {
		// 真下
		if self.Fields[a][x].ShipID != 0 {
			return true
		}

		d := x + 1
		e := x - 1

		if d < len(self.Fields[a]) {
			// 右下
			if self.Fields[a][d].ShipID != 0 {
				return true
			}
		}

		if e >= 0 {
			// 左下
			if self.Fields[a][e].ShipID != 0 {
				return true
			}
		}

	}

	if b >= 0 {
		// 真上
		if self.Fields[b][x].ShipID != 0 {
			return true
		}

		d := x + 1
		e := x - 1

		if d < len(self.Fields[b]) {
			// 右上
			if self.Fields[b][d].ShipID != 0 {
				return true
			}
		}

		if e >= 0 {
			// 左上
			if self.Fields[b][e].ShipID != 0 {
				return true
			}
		}

	}

	d := x + 1
	e := x - 1

	if d < len(self.Fields[y]) {
		// 右
		if self.Fields[y][d].ShipID != 0 {
			return true
		}
	}

	if e >= 0 {
		// 左
		if self.Fields[y][e].ShipID != 0 {
			return true
		}
	}

	return false
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

func (self *Room) GetUserFromConn(conn *websocket.Conn) (string, *User, error) {
	for userID, user := range self.Users {
		if user.Conn == conn {
			return userID, user, nil
		}
	}
	return "", nil, fmt.Errorf("NOT_FOUND")
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

func (self *Room) GetDestroyShipSize(enemy *User, target *Field) int {

	if target.ShipID == 0 {
		return 0
	}

	size := 0
	for _, line := range enemy.Fields {
		for _, field := range line {
			if field.ShipID == target.ShipID {
				// 攻撃受けてないパーツ
				if field.HitType == 0 {
					return 0
				} else {
					size = size + 1
				}
			}
		}
	}
	return size
}
