package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/tomyhero/battleship-game/matching/server"
	"github.com/tomyhero/battleship-game/utils"
)

var flagValue struct {
	ConfigPath string
	Port       int
}

func init() {
	flag.StringVar(&flagValue.ConfigPath, "config", "./etc/config/html5-example.toml", "set config path")
	flag.IntVar(&flagValue.Port, "port", 8080, "port")
	flag.Parse()
}

func main() {

	config, err := utils.NewConfigFromFile(flagValue.ConfigPath)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("Fail to load config")
	}

	log.WithFields(log.Fields{"config": config}).Info("Loaded Config")

	server := server.NewServer()
	log.WithFields(log.Fields{"port": flagValue.Port}).Info("Start Listen And Serve")
	server.ListenAndServe(flagValue.Port)
}
