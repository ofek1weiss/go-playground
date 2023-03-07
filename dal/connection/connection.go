package connection

import (
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

const fileName = "tasks.db"

func Connect() (*bolt.DB, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(home, fileName)
	return bolt.Open(filePath, 0600, nil)
}

func Close(db *bolt.DB) {
	db.Close()
}

func Context(f func(*bolt.DB) error) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer Close(db)
	return f(db)
}
