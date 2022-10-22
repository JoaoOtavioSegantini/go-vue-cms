package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/service"
)

type PostController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	service service.PostService
}

func New(service service.PostService) PostController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	c.service.FindAll(ctx)

}

func (c *controller) Save(ctx *gin.Context) {
	c.service.Save(ctx)
}

func (c *controller) Update(ctx *gin.Context) {

	c.service.Update(ctx)
}

func (c *controller) FindById(ctx *gin.Context) {
	c.service.FindById(ctx)
}

func (c *controller) Delete(ctx *gin.Context) {

	c.service.Delete(ctx)
}
