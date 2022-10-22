package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
	"github.com/joaotavioos/cms-server/service"
)

func RegisterUser(context *gin.Context) {
	var user entity.UserMysql
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

type UserController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUsers(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) FindAll(ctx *gin.Context) {
	c.service.FindAll(ctx)

}

func (c *userController) Save(ctx *gin.Context) {
	c.service.Save(ctx)
}

func (c *userController) Update(ctx *gin.Context) {

	c.service.Update(ctx)
}

func (c *userController) FindById(ctx *gin.Context) {
	c.service.FindById(ctx)
}

func (c *userController) Delete(ctx *gin.Context) {

	c.service.Delete(ctx)
}
