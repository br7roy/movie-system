package service

import (
	"database/sql"
	"fmt"
	"movie-system/app/kit"
	"movie-system/app/model"
	_ "movie-system/app/model"
)

func GetAllMovies() []model.Movie {
	db, err := sql.Open("sqlite3", "./movie.db")
	kit.IfNull(err)

	sql_table := `

      select * from movie ;
    `
	query, err := db.Query(sql_table)

	kit.IfNull(err)

	var mname string
	var mdesc string
	var mpath string
	var movies []model.Movie

	for query.Next() {
		err := query.Scan(&mname, &mdesc, &mpath)

		kit.IfNull(err)
		fmt.Println(mname)
		fmt.Println(mdesc)
		fmt.Println(mpath)
		movie := model.Movie{Mpath: mpath, Mdesc: mdesc, Mname: mname}
		movies = append(movies, movie)
	}

	_ = query.Close()

	for _, movie := range movies {
		fmt.Println(movie)
	}
	return movies

}
