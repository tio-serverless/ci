package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	path := os.Getenv("TIO_TRCI_CONF")
	if path == "" {
		path = "trci.toml"
	}

	c, err := NewConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range c.Repos.Branch {
		for _, n := range c.Repos.Name {
			log.Debugf("Trigger %s:%s", n, b)
			if err := trigger(b, n, c.Token); err != nil {
				log.Errorf("Trigger %s:%s Error. %s", n, b, err.Error())
			}
		}

	}

}
