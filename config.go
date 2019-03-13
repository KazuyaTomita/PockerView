package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server ServerConfig `toml:"server"`
	Engines MultiEnginesConfig `toml:"multiEngines"`
	Cli CliConfig `toml:"cli"`
}

type ServerConfig struct {
	Enable bool `toml:"enable"`
	Ip  string `toml:"ip"`
	Port  string `toml:"port"`
	Path string `toml:"path"`
}

type MultiEnginesConfig struct {
	Enable bool `toml:"enable"`
	Number int `toml:"number"`
	Paths []string `toml:"paths"`
}

type CliConfig struct {
	Enable bool `toml:"enable"`
	Path string `toml:"path"`
}

func ReadConfig(config *Config) {
	// read config file
	_, parseErr := toml.DecodeFile("config.toml", config)
	if parseErr != nil {
		panic(parseErr)
	}
	// check whether we can decide which mode is used
	if isMultipleModesSelected(config) {
		panic("not sure of which mode is used. please check config. Only single enable flag should be true.")
	}
}

func isMultipleModesSelected(config *Config) bool {
	return (config.Server.Enable && config.Engines.Enable) ||
		(config.Server.Enable && config.Cli.Enable) ||
		(config.Cli.Enable && config.Engines.Enable)
}
