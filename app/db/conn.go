package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"movie-system/app/kit"
	_ "movie-system/app/kit"
)

var Db *sql.DB

func init() {
	var err error
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", kit.AppConfig.Database.UserName, kit.AppConfig.Database.Password, "tcp", kit.AppConfig.Database.Url, kit.AppConfig.Database.Port, kit.AppConfig.Database.StoreHouse)
	Db, err = sql.Open(kit.AppConfig.Database.Type, conn)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
