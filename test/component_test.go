package test

import (
	"moviex/app/db"
	"moviex/app/kit"
	"moviex/app/model"
	"testing"
)

func TestName(t *testing.T) {
	query, err := db.Db.Query("select * from `movie`")
	kit.IfNull(err)

	var id int
	var name string
	var desc string
	var path string
	var movies []model.Movie

	for query.Next() {
		err := query.Scan(&id, &name, &desc, &path)

		kit.IfNull(err)

		movie := model.Movie{Id: id, Path: path, Desc: desc, Name: name}
		movies = append(movies, movie)
	}
}
