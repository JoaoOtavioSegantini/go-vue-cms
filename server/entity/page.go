package entity

import "gorm.io/gorm"

type Page struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PageMysql struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}
