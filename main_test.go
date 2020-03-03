package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	_ "movie-system/app/kit"
	"movie-system/app/service"
	"testing"
)

func TestA(t *testing.T) {
	movies := service.GetAllMovies()
	for _, movie := range movies {
		fmt.Println(movie)
	}

}
