package bootstrap

import (
	"github.com/go-pg/pg"
	"github.com/josephniel/go-api/app/config"
	"github.com/josephniel/go-api/app/models"
)

// SetupDB sets up the database connection to be used by the application
func SetupDB(conf *config.Config) {
	models.DB = pg.Connect(&pg.Options{
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Database: conf.DB.Database,
	})
}

// CloseDB is a wrapper for the pg.DB.Close function which closes the database connection
func CloseDB() {
	models.DB.Close()
}
