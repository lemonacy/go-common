package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/lemonacy/go-common/db/mysql"
	"github.com/lemonacy/go-common/hoo"
	"github.com/lemonacy/go-common/web"
)

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	CreatedTime  time.Time `json:"createdTime"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedTime time.Time `json:"modifiedTime"`
	ModifiedBy   string    `json:"modifiedBy"`
}

func main() {
	web.Router.POST("/insertUser", func(c *gin.Context) {
		rs, err := db.MySQL.Exec("insert into t_user(name, age) values(?, ?)", "Asran", 36)
		hoo.PanicOnErr(err)
		n, err := rs.LastInsertId()
		hoo.PanicOnErr(err)
		c.String(http.StatusOK, "New record id is %d.", n)
	})

	web.Router.GET("/getUser", func(c *gin.Context) {
		id, b := c.GetQuery("id")
		if !b {
			panic("user id must be specified.")
		}
		row := db.MySQL.QueryRow("select * from t_user where id = ?", id)
		user := User{}
		err := row.Scan(&user.ID, &user.Name, &user.Age, &user.CreatedTime, &user.CreatedBy, &user.ModifiedTime, &user.ModifiedBy)
		hoo.PanicOnErr(err)
		c.JSON(http.StatusOK, user)
	})

	web.Router.PUT("/updateUser", func(c *gin.Context) {
		id, b := c.GetQuery("id")
		if !b {
			panic("user id must be specified.")
		}
		params := map[string]interface{}{}
		err := c.BindJSON(&params)
		hoo.PanicOnErr(err)
		rs, err := db.MySQL.Exec("update t_user set age = ? where id = ?", params["age"], id)
		hoo.PanicOnErr(err)
		n, err := rs.RowsAffected()
		hoo.PanicOnErr(err)
		if n != 1 {
			panic("update user failed.")
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success"})
	})

	web.Router.DELETE("/deleteUser", func(c *gin.Context) {
		id, b := c.GetQuery("id")
		if !b {
			panic("user id must be specified.")
		}
		rs, err := db.MySQL.Exec("delete from t_user where id = ?", id)
		hoo.PanicOnErr(err)
		n, err := rs.RowsAffected()
		hoo.PanicOnErr(err)
		if n != 1 {
			panic("delete user failed.")
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success"})
	})

	web.Router.GET("/getAllUsers", func(c *gin.Context) {
		result := db.GenericQuery("select * from t_user")
		c.JSON(http.StatusOK, result)
	})

	web.Run()
}
