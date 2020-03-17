package test

import (
	"fmt"
	"io/ioutil"
	"movie-system/app/db"
	"movie-system/app/model"
	"testing"
	"unsafe"
)

func TestInsert(t *testing.T) {
	u := model.User{Email: "gorm@email", Password: "123123"}
	db.Db.Create(&u)

}

func TestDelete(t *testing.T) {
	u := model.User{
		//ID: 29,
	}
	where := db.Db.Where("id = ?", 31).Delete(u)
	fmt.Println(where.RowsAffected)

}

func TestUpdate(t *testing.T) {

	user := model.User{
		ID:     1,
		Avatar: "https://tlasdf",
	}

	update := db.Db.Model(&user).Update("avatar", user.Avatar)
	affected := update.RowsAffected
	fmt.Println(affected)

}

func TestFindOne(t *testing.T) {
	user := model.User{}

	db.Db.Find(&user).Where("id = ?", 1)
	fmt.Println(user)

}

func TestFindOne2(t *testing.T) {
	user := model.User{
		ID:    30,
		Email: "123@gmail.com",
		//Password: "123",
	}
	db.Db.Find(&user)
	fmt.Println(user)
}

func TestFinUpdate(t *testing.T) {
	user := model.User{}

	db.Db.Model(&user).Where("id = ?", 1).Update("avatar", "sadfasdfoijo")
	fmt.Println(user)

}

func TestFindAllMovie(t *testing.T) {

	var movies []model.Movie
	db.Db.Find(&movies)
	fmt.Println(movies)
}

func TestReadFile(t *testing.T) {

	b, _ := ioutil.ReadFile("c:\\users\\tak\\desktop\\TIM图片20200317150736.png")
	s := *(*string)(unsafe.Pointer(&b))
	println(s)

}
