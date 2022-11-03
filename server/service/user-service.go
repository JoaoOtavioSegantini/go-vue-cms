package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
)

type UserService interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userService struct {
	users []entity.UserMysql
}

func NewUsers() UserService {
	return &userService{}
}

func (service *userService) Save(ctx *gin.Context) {
	var user entity.UserMysql
	ctx.BindJSON(&user)
	database.Instance.Create(&user)

	ctx.JSON(http.StatusCreated, user)
}

func (service *userService) Update(ctx *gin.Context) {
	var userData entity.UserMysql

	id := ctx.Param("id")
	ctx.BindJSON(&userData)
	database.Instance.Model(&userData).Where("ID = ?", id).Updates(&userData).Find(&userData)
	ctx.JSON(http.StatusOK, &userData)
}

func (service *userService) FindAll(ctx *gin.Context) {
	var users []entity.UserMysql

	database.Instance.Select("ID", "Name", "Username", "Email", "CreatedAt", "UpdatedAt", "DeletedAt").Find(&users)
	ctx.JSON(http.StatusOK, users)

}

func (service *userService) FindById(ctx *gin.Context) {
	var user entity.UserMysql

	id := ctx.Param("id")
	database.Instance.Model(&user).Where("ID = ?", id).Find(&user)
	ctx.JSON(http.StatusOK, user)

}

func (service *userService) Delete(ctx *gin.Context) {
	var user entity.UserMysql

	id := ctx.Param("id")

	database.Instance.Model(&user).Where("ID = ?", id).Delete(&user)
	ctx.JSON(http.StatusNoContent, nil)
}
