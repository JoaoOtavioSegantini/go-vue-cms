package entity

import "gorm.io/gorm"

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}

type PostMysql struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}
