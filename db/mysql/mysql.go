package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lemonacy/go-common/config"
	"github.com/lemonacy/go-common/hoo"
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

func GenericQuery(sql string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	rs, err := MySQL.Query(sql)
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
			switch val.(type) {
			case int64:
				data[name] = val.(int64)
			case string:
				data[name] = val.(string)
			case time.Time:
				data[name] = val.(time.Time)
			case []uint8:
				data[name] = string(val.([]uint8))
			default:
				data[name] = val
			}
		}
		result = append(result, data)
	}
	return result
}
