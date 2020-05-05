package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/lemonacy/go-common/config"
)

var Postgres *sql.DB

func init() {
	var err error
	Postgres, err = sql.Open("postgres", config.Viper.GetString("database.postgres.url"))
	if err != nil {
		log.Fatal(err)
	}

	Postgres.Ping()
}
