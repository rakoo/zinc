package main

import (
	"flag"
	"log"

	couch "github.com/dustin/go-couch"
)

var (
	flagCmd    = flag.String("cmd", "", "The command to launch")
	flagSource = flag.String("source", "", "The source of the watched dir")
	flagPath   = flag.String("path", "", "The path of the file/dir")
)

type FileRecord struct {
	couch.IdAndRev

	Root string
	Path string
	Sha1 string
}

func main() {
	flag.Parse()

	if *flagCmd == "" {
		log.Fatal("Need a cmd to run")
	}

	if *flagSource == "" {
		log.Fatal("Need source to work")
	}

	db, cerr := couch.NewDatabase("localhost", "5984", "zinc-pub")
	if cerr != nil {
		log.Fatal("Error when creating database: ", cerr)
	}

	var err error

	switch *flagCmd {
	case "onStartup":
		err = startup(db, *flagSource)
	case "onCreate":
		mustNotEmpty(*flagPath)
		err = create(db, *flagSource, *flagPath)

	case "onModify":
		mustNotEmpty(*flagPath)
	case "onMove":
		mustNotEmpty(*flagPath)
	case "onDelete":
		mustNotEmpty(*flagPath)

	default:
		log.Println("Unknown command:", *flagCmd)
	}

	if err != nil {
		log.Fatal("Error with action")
	}

}

func mustNotEmpty(fl string) {
	if fl == "" {
		log.Fatal("Path is empty!")
	}
}
