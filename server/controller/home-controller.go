package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
)

func Home(context *gin.Context) {
	var page entity.PageMysql
	var menus []entity.PageMysql
	var posts []entity.PostMysql
	var total int64

	offset := context.Request.URL.Query().Get("page")

	if offset == "" {
		offset = "0"
	}

	pagination, err := strconv.Atoi(offset)

	if err != nil {
		context.Abort()
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	database.Instance.Model(&page).First(&page)
	database.Instance.Find(&menus)
	database.Instance.Model(&posts).Count(&total).Offset(pagination).Limit(10).Find(&posts)

	context.HTML(http.StatusOK, "page.html", gin.H{"page": page, "menus": menus, "posts": posts})

}

func View(context *gin.Context) {
	var page entity.PageMysql

	slug := context.Param("slug")
	database.Instance.Model(&page).Where("Slug = ?", slug).Find(&page)
	var menus []entity.PageMysql

	database.Instance.Find(&menus)
	context.HTML(http.StatusOK, "view.html", gin.H{"page": page, "menus": menus})
}

func Show(context *gin.Context) {
	var menus []entity.PageMysql
	var page entity.PageMysql
	var post entity.PostMysql

	database.Instance.Find(&menus)

	slug := context.Param("slug")
	database.Instance.Model(&post).Where("Slug = ?", slug).Find(&post)
	context.HTML(http.StatusOK, "articles.html", gin.H{"page": page, "menus": menus, "post": post})

}
