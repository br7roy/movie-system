package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//数据库连接信息
const (
	USERNAME = "movie-system"
	PASSWORD = "ZaW6rjRcee26ijEc"
	NETWORK  = "tcp"
	SERVER   = "158.51.124.231"
	PORT     = 3306
	DATABASE = "movie-system"
)

var Db *sql.DB

func init() {
	var err error
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	Db, err = sql.Open("mysql", conn)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
