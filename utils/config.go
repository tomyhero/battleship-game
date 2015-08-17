package utils

import (
	"github.com/BurntSushi/toml"
	"path"
)

// Config struct hold your configuration values
type Config struct {
	GameServer        GameServerSection        `toml:"game_server"`
	MatchingServer    MatchingServerSection    `toml:"matching_server"`
	HTML5ClientServer HTML5ClientServerSection `toml:"html5_client_server"`
}

type MatchingServerSection struct {
	Endpoint string `toml:"endpoint"`
}

type GameServerSection struct {
	Endpoint string `toml:"endpoint"`
}

type HTML5ClientServerSection struct {
	AssetsPath string `toml:"assets_path"`
}

// NewConfigFromFile create Config instance. you need to set toml formated config location path.
func NewConfigFromFile(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	// TODO config data validate

	return &config, nil
}

func (config *Config) TemplatePath(template string) string {
	assetsPath := path.Join(config.HTML5ClientServer.AssetsPath, "template", template+".tpl")
	return assetsPath
}
