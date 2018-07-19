package bootstrap

import (
	"github.com/go-pg/pg"
)

// DB is the wrapper structure for database related entities
type DB struct {
	User         string
	Password     string
	DatabaseName string
	DatabaseObj  **pg.DB
}

// Start sets up the database connection to be used by the application
func (db *DB) Start() {
	*db.DatabaseObj = pg.Connect(&pg.Options{
		User:     db.User,
		Password: db.Password,
		Database: db.DatabaseName,
	})
}

// Close is a wrapper for the pg.DB.Close function which closes the database connection
func (db *DB) Close() {
	(*db.DatabaseObj).Close()
}
