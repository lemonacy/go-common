package db

import (
	"github.com/lemonacy/go-common/config"
	"github.com/lemonacy/go-common/hoo"
	"gopkg.in/mgo.v2"
)

var Mongo *mgo.Session

func init() {
	var err error
	Mongo, err = mgo.Dial(config.Viper.GetString("database.mongo.url"))
	hoo.PanicOnErr(err)
	hoo.PanicOnErr(Mongo.Ping())
}
