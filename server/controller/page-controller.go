package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/service"
)

type PageController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type pageController struct {
	service service.PageService
}

func NewPage(service service.PageService) PageController {
	return &pageController{
		service: service,
	}
}

func (c *pageController) FindAll(ctx *gin.Context) {
	c.service.FindAll(ctx)

}

func (c *pageController) Save(ctx *gin.Context) {
	c.service.Save(ctx)
}

func (c *pageController) Update(ctx *gin.Context) {

	c.service.Update(ctx)
}

func (c *pageController) FindById(ctx *gin.Context) {
	c.service.FindById(ctx)
}

func (c *pageController) Delete(ctx *gin.Context) {

	c.service.Delete(ctx)
}
