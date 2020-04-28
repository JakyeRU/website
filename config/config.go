package config

import "github.com/BurntSushi/toml"

var Conf Config

func LoadConfig() {
	if _, err := toml.DecodeFile("config.toml", &Conf); err != nil {
		panic(err)
	}
}
