package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/lemonacy/go-common/config"
	"github.com/lemonacy/go-common/hoo"
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

func GenericQuery(sql string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	rs, err := Postgres.Query(sql)
	hoo.PanicOnErr(err)
	cols, err := rs.Columns()
	hoo.PanicOnErr(err)
	for rs.Next() {
		vals := make([]interface{}, len(cols))
		ptrs := make([]interface{}, len(cols))
		for idx := range vals {
			ptrs[idx] = &vals[idx]
		}
		err := rs.Scan(ptrs...)
		hoo.PanicOnErr(err)
		data := make(map[string]interface{})
		for idx, name := range cols {
			val := *(ptrs[idx].(*interface{}))
			// switch val.(type) {
			// case []uint8:
			// 	data[name] = string(val.([]uint8))
			// default:
			data[name] = val
			// }
		}
		result = append(result, data)
	}
	return result
}
