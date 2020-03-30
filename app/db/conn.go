package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"moviex/app/kit"
	_ "moviex/app/kit"
)

//var Db *sql.DB
var Db *gorm.DB

func init() {
	var err error
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", kit.AppConfig.Database.UserName, kit.AppConfig.Database.Password, "tcp", kit.AppConfig.Database.Url, kit.AppConfig.Database.Port, kit.AppConfig.Database.StoreHouse)

	Db, err = gorm.Open(kit.AppConfig.Database.Type, conn)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SingularTable(true)
	//Db.SetMaxOpenConns(10)
	//Db.SetMaxIdleConns(10)
}
