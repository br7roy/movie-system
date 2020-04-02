package model

import (
	"moviex/app/db"
)

type User struct {
	ID       int    `form:"id",gorm:"primary_key"`
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
	//PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
	//PasswordAgain string `form:"password-again"`
	Avatar string
}

var UserInfo = make(map[int]User)

func Save(user *User) int64 {

	model := User{Email: user.Email, Password: user.Password}

	//result:= db.Db.Create("insert into user (email,pwd) values(?,?)", user.Email, user.Password)
	//id, err := result.LastInsertId()
	//if err != nil {
	//	log.Panicln("user insert id error", err.Error())
	//}

	db.Db.Create(&model)
	id := model.ID

	return int64(id)

}

func (user *User) QueryByEmail() User {
	//u := User{}
	db.Db.Find(&user)
	//row := db.Db.QueryRow("select * from user where email = ?;", user.Email)
	//e := row.Scan(&u.ID, &u.Email, &u.Password, &u.Avatar)
	//if e != nil {
	//	log.Panicln(e)
	//}
	return *user
}

func (user *User) QueryById(id int) (User, error) {
	//u := User{}
	//first := db.Db.First(&u, id)
	//e := first.Error
	//row := db.Db.QueryRow("select * from user where id = ?;", id)
	//e := row.Scan(&u.ID, &u.Email, &u.Password, &u.Avatar)
	//if e != nil {
	//	log.Panicln(e)
	//}
	e := db.Db.Find(user).Where("id=?", id).Error

	return *user, e
}

func (user *User) Update(id int) int64 {

	// Example:
	//  updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
	//  ...
	//  tx, err := db.Begin()
	//  ...
	//  res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)
	//prepare, err := db.Db.Prepare("update user set avatar = ? where id = ?;")
	//kit.IfNull(err)
	//tx, err := db.Db.Begin()
	//res, err := tx.Stmt(prepare).Exec(user.Avatar, user.ID)
	//affected, err := res.RowsAffected()
	//kit.IfNull(err)

	update := db.Db.Model(&user).Where("id=?", user.ID).Update("avatar", user.Avatar)
	affected := update.RowsAffected

	return affected

}
