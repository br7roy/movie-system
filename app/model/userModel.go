package model

import (
	"database/sql"
	"log"
	"movie-system/app/db"
	"movie-system/app/kit"
)

type UserModel struct {
	Id       int    `form:"id"`
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
	//PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
	PasswordAgain string `form:"password-again"`
	Avatar        sql.NullString
}

func Save(user *UserModel) int64 {

	result, e := db.Db.Exec("insert into user (email,pwd) values(?,?)", user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}

	return id

}

func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := db.Db.QueryRow("select * from user where email = ?;", user.Email)
	e := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Panicln(e)
	}
	return u
}

func (user *UserModel) QueryById(id int) (UserModel, error) {
	u := UserModel{}
	row := db.Db.QueryRow("select * from user where id = ?;", id)
	e := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Panicln(e)
	}
	return u, e
}

func (user *UserModel) Update(id int) int64 {

	// Example:
	//  updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
	//  ...
	//  tx, err := db.Begin()
	//  ...
	//  res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)

	prepare, err := db.Db.Prepare("update user set avatar = ? where id = ?;")
	kit.IfNull(err)
	tx, err := db.Db.Begin()
	res, err := tx.Stmt(prepare).Exec(user.Avatar, user.Id)
	affected, err := res.RowsAffected()
	kit.IfNull(err)

	return affected

}
