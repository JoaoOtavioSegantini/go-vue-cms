package entity

import (
	"html/template"

	"gorm.io/gorm"
)

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}

type PostMysql struct {
	gorm.Model
	Title string        `json:"title" gorm:"unique"`
	Body  template.HTML `json:"body"`
	Slug  string        `json:"slug"`
}
