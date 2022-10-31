package entity

import (
	"html/template"

	"gorm.io/gorm"
)

type Page struct {
	//	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`

	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PageMysql struct {
	gorm.Model
	Title string        `json:"title" gorm:"unique"`
	Body  template.HTML `json:"body"`
	Slug  string        `json:"slug"`
}
