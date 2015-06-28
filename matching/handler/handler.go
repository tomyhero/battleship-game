package handler

import (
	"fmt"
	"github.com/tomyhero/battleship-game/matching/data"
)

type Handler struct {
}

func (self Handler) Search(data data.Interface) {
	fmt.Println(data, "Search")
}
