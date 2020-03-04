package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"movie-system/app/kit"
	"movie-system/app/model"
	"net/http"
)

func MovieIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Movie coming.",
		"text":  "Hello Gin",
	})
}

func ShowAllMovies(context *gin.Context) {
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

	if gin.Mode() == gin.DebugMode {
		for _, movie := range movies {
			fmt.Println(movie)
		}
	}

	context.HTML(http.StatusOK, "movie_index.tmpl", gin.H{
		"movies": movies,
	})

}
