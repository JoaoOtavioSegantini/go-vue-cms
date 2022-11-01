package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
)

func Home(context *gin.Context) {

	var (
		page               entity.PageMysql
		menus              []entity.PageMysql
		posts              []entity.PostMysql
		total              int64
		prevPage, nextPage string
	)

	const limit = 10

	pageStr := context.DefaultQuery("page", "1")
	pagination, err := strconv.Atoi(pageStr)

	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	database.Instance.Model(&page).First(&page)
	database.Instance.Find(&menus)
	database.Instance.Model(&posts).Count(&total)

	pageCount := int(math.Ceil(float64(total) / float64(limit)))

	if pageCount == 0 {
		pageCount = 1
	}

	if pagination < 1 || pagination > pageCount {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	offset := (pagination - 1) * limit

	database.Instance.Model(&posts).Offset(offset).Limit(limit).Find(&posts)

	if pagination > 1 {
		prevPage = fmt.Sprintf("%d", pagination-1)
	}
	if pagination < pageCount {
		nextPage = fmt.Sprintf("%d", pagination+1)
	}

	pages := make([]int, pageCount)

	for i := 0; i < pageCount; i++ {
		pages[i] = i + 1
	}

	context.HTML(http.StatusOK, "page.html", gin.H{
		"page":       page,
		"menus":      menus,
		"posts":      posts,
		"pageCount":  pageCount,
		"pagination": pagination,
		"prevPage":   prevPage,
		"nextPage":   nextPage,
		"pages":      pages,
	})
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
