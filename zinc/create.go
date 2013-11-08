package main

import (
	couch "github.com/dustin/go-couch"
	"log"
	"net/url"
	"regexp"
)

func create(db couch.Database, source, path string) error {
  matched, err := regexp.MatchString("/$", path)
  if err != nil {
    return err
  }
  if matched {
    return nil
  }

	protected := url.QueryEscape(path)

	var fr FileRecord
	err = db.Retrieve(protected, fr)

	if err != nil {
		if err.(*couch.HttpError).Status != 404 {
			log.Println("Error when creating file in couchdb: ", err)
		}
	}

	if fr.Id != "" {
		log.Println("File record already existing")
	}

	fr = FileRecord{
		Root: source,
		Path: path,
		Sha1: "",
	}
  fr.Id = protected

	_, _, err = db.Insert(fr)

	if err != nil {
		log.Println("Error when putting file record: ", err)
		return err
	}

	return nil
}
