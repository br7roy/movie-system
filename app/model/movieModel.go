package model

type Movie struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Path string `json:"path"`
}
