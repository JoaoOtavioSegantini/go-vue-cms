package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
)

type PostService interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type postService struct {
	posts []entity.Post
}

func New() PostService {
	return &postService{}
}

func (service *postService) Save(ctx *gin.Context) {
	var post entity.PostMysql
	ctx.BindJSON(&post)
	text := slug.Make(post.Title)
	post.Slug = text
	database.Instance.Create(&post)

	ctx.JSON(http.StatusCreated, post)
}

func (service *postService) Update(ctx *gin.Context) {
	var postData entity.PostMysql

	id := ctx.Param("id")
	ctx.BindJSON(&postData)
	postData.Slug = slug.Make(postData.Title)
	database.Instance.Model(&postData).Where("ID = ?", id).Updates(&postData)
}

func (service *postService) FindAll(ctx *gin.Context) {
	var posts []entity.PostMysql

	database.Instance.Find(&posts)
	ctx.JSON(http.StatusOK, posts)

}

func (service *postService) FindById(ctx *gin.Context) {
	var post entity.PostMysql

	id := ctx.Param("id")
	database.Instance.Model(&post).Where("ID = ?", id).Find(&post)
	ctx.JSON(http.StatusOK, post)

}

func (service *postService) Delete(ctx *gin.Context) {
	var post entity.PostMysql

	id := ctx.Param("id")

	database.Instance.Model(&post).Where("ID = ?", id).Delete(&post)
	ctx.JSON(http.StatusNoContent, nil)
}
