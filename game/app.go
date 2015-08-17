package main

import (
	"flag"
	"fmt"
	"github.com/tomyhero/battleship-game/game/server"
	"github.com/tomyhero/battleship-game/utils"
)

var flagValue struct {
	ConfigPath string
	Port       int
}

func init() {
	flag.StringVar(&flagValue.ConfigPath, "config", "./etc/config/html5-example.toml", "set config path")
	flag.IntVar(&flagValue.Port, "port", 9090, "port")
	flag.Parse()
}

func main() {

	config, err := utils.NewConfigFromFile(flagValue.ConfigPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Loaded Config", config)

	server := server.NewServer()
	server.ListenAndServe(flagValue.Port)
}
