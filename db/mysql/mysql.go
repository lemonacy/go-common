package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lemonacy/go-common/config"
)

var MySQL *sql.DB

func init() {
	var err error
	MySQL, err = sql.Open("mysql", config.Viper.GetString("database.mysql.url"))
	if err != nil {
		log.Fatal(err)
	}

	MySQL.Ping()
}
