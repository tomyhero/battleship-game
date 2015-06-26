package main

import (
	"github.com/tomyhero/submarine-game/matching/server"
)

func main() {
	server := server.Server{}
	server.ListenAndServe()
}
