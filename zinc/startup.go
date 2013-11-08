package main

import (
	"errors"
	"log"

	couch "github.com/dustin/go-couch"
)

func startup(db couch.Database, source string) error {
	log.Println("Running onStartup in", *flagSource)

	if db.Exists() {
		return nil
	} else {
		return errors.New("Database doesn't exist")
	}
}
