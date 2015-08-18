package data

type Room struct {
	Status      int
	Users       map[string]User
	GameSetting GameSetting
}

func NewRoom() Room {
	room := Room{Users: map[string]User{}, GameSetting: NewGameSetting()}
	return room
}

type GameSetting struct {
	GridX int
	GridY int
}

func NewGameSetting() GameSetting {
	return GameSetting{GridX: 16, GridY: 16}
}

type User struct {
}
