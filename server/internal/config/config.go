package config

import "database/sql"

type Config struct {
	Port string
	Db   *sql.DB
}
