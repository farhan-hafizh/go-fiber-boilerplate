package config

import (
	"os"
)

// DB holds the DB configuration
type DB struct {
	Prefix   string
	Host     string
	Name     string
	Username string
	Password string
}

var db = &DB{}

// DBCfg returns the default DB configuration
func DBCfg() *DB {
	return db
}

// LoadDBCfg loads DB configuration
func LoadDBConfig() {
	db.Prefix = os.Getenv("DB_PREFIX")
	db.Host = os.Getenv("DB_HOST")
	db.Username = os.Getenv("DB_USERNAME")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Name = os.Getenv("DB_NAME")
}
