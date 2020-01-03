package main

import (
	"strings"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

var _VERSION_, _BRANCH_ string

type config struct {
	Log   string `toml:"log"`
	Token string `toml:"token"`
	Repos []struct {
		Branch []string `toml:"branch"`
		Name   []string `toml:"name"`
	} `toml:"repos"`
}

func NewConfig(path string) (config, error) {
	c := config{}
	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return c, err
	}

	enableDebug(c)
	output(c)

	return c, nil
}

func enableDebug(c config) {
	switch strings.ToLower(c.Log) {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}

func output(c config) {
	log.Infof("TrCi Config. Version: %s Branch: %s", _VERSION_, _BRANCH_)
	log.Infof("Debug: %s", c.Log)
	log.Info("Repos:")
	for _, r := range c.Repos {
		log.Info("  Branch:")
		for _, b := range r.Branch {
			log.Infof("    %s", b)
		}
		log.Info("  Names:")
		for _, n := range r.Name {
			log.Infof("    %s", n)
		}
	}

	log.Info("")
}
