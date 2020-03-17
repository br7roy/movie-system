package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie-system/app/db"
	"movie-system/app/model"
	"net/http"
)

func MovieIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":    "Movie coming.",
		"navTitle": "Hello World",
	})
}

func ShowAllMovies(context *gin.Context) {

	//sql_table := "select * from movie"

	//query, err := db.Db.Query(sql_table)

	//kit.IfNull(err)
	/*	for query.Next() {
			err := query.Scan(&id, &name, &desc, &path)

			kit.IfNull(err)

			movie := model.Movie{ID: id, Path: path, Desc: desc, Name: name}
			movies = append(movies, movie)
		}
		var id int
		var name string
		var desc string
		var path string
			_ = query.Close()
	*/
	var movies []model.Movie
	db.Db.Find(&movies)

	if gin.Mode() == gin.DebugMode {
		for _, movie := range movies {
			fmt.Println(movie)
		}
	}

	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"movies": movies,
	})

}
