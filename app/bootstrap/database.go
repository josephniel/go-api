package bootstrap

import (
	"github.com/go-pg/pg"
)

// DB is the wrapper structure for database related entities
type DB struct {
	User         string
	Password     string
	DatabaseName string
	Database     **pg.DB
}

// SetupDB sets up the database connection to be used by the application
func (db *DB) SetupDB() {
	*db.Database = pg.Connect(&pg.Options{
		User:     db.User,
		Password: db.Password,
		Database: db.DatabaseName,
	})
}

// CloseDB is a wrapper for the pg.DB.Close function which closes the database connection
func (db *DB) CloseDB() {
	(*db.Database).Close()
}
