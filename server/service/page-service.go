package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
)

type PageService interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func (service *pageService) Save(ctx *gin.Context) {
	var page entity.PageMysql
	ctx.BindJSON(&page)
	text := slug.Make(page.Title)
	page.Slug = text
	database.Instance.Create(&page)

	ctx.JSON(http.StatusCreated, page)
}

type pageService struct {
	page []entity.PageMysql
}

func NewPage() PageService {
	return &pageService{}
}

func (service *pageService) Update(ctx *gin.Context) {
	var pageData entity.PageMysql

	id := ctx.Param("id")
	ctx.BindJSON(&pageData)
	pageData.Slug = slug.Make(pageData.Title)
	database.Instance.Model(&pageData).Where("ID = ?", id).Updates(&pageData).Find(&pageData)
	ctx.JSON(http.StatusOK, pageData)

}

func (service *pageService) FindAll(ctx *gin.Context) {
	var pages []entity.PageMysql

	database.Instance.Find(&pages)
	ctx.JSON(http.StatusOK, pages)

}

func (service *pageService) FindById(ctx *gin.Context) {
	var page entity.PageMysql

	id := ctx.Param("id")
	database.Instance.Model(&page).Where("ID = ?", id).Find(&page)
	ctx.JSON(http.StatusOK, page)

}

func (service *pageService) Delete(ctx *gin.Context) {
	var page entity.PageMysql

	id := ctx.Param("id")

	database.Instance.Model(&page).Where("ID = ?", id).Delete(&page)
	ctx.JSON(http.StatusNoContent, nil)
}
