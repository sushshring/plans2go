package db

import (
	"log"
	"os"

	"github.com/boltdb/bolt"
)

// Config is the configuration for the database
type Config struct {
	filename string
	mode     os.FileMode
	options  *bolt.Options
}

// Init initializes the database
func Init(config Config) *bolt.DB {
	db, err := bolt.Open(config.filename, config.mode, config.options)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
