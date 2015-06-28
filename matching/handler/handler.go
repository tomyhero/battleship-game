package handler

import (
	"fmt"
	"github.com/tomyhero/battleship-game/matching/data"
)

type Handler struct {
}

func (self Handler) Search(data data.Interface) {
	fmt.Println(data, "Search")

	// found

	// generate MatchID
	// modify myself and enemy user matchID if fail then make it not found
	// tell to enemy
	// return game id to client
	// dissconnect

	//  not found

	// do nothing.

}

func (seslf Handler) Found(data data.Interface) {
	fmt.Println(data, "Found")

}
