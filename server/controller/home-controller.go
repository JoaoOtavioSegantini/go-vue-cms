package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
)

func Home(context *gin.Context) {
	var page entity.PageMysql
	var menus []entity.PageMysql

	database.Instance.Model(&page).First(&page)
	database.Instance.Find(&menus)

	context.HTML(http.StatusOK, "page.html", gin.H{"page": page, "menus": menus})

}

func View(context *gin.Context) {
	var page entity.PageMysql

	slug := context.Param("slug")
	database.Instance.Model(&page).Where("Slug = ?", slug).Find(&page)

	context.HTML(http.StatusOK, "view.html", gin.H{"page": page})
}
