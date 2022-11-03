package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
	"github.com/joaotavioos/cms-server/routes"
	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	//	database.Connect("root:root@tcp(localhost:3306)/cms-database?parseTime=true")
	database.Connect("")

	database.Migrate()

	var (
		page, page1, page2, page3 entity.PageMysql
		post, post1, post2, post3 entity.PostMysql
	)

	page.Title = "Home"
	page.Body = "<p>&nbsp;</p><figure class=\"image\"><img src=\"http://localhost:8000/uploads/undraw_writer_q06d.png\"></figure><p><strong>Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis.</strong></p>"
	page.Slug = slug.Make(page.Title)
	page1.Title = "Links"
	page1.Slug = slug.Make(page1.Title)
	page2.Title = "Sobre nós"
	page2.Slug = slug.Make(page2.Title)
	page3.Title = "Fale conosco"
	page3.Slug = slug.Make(page3.Title)

	post.Title = "A simple test title for first post"
	post.Slug = slug.Make(post.Title)
	post1.Title = "A simple test title for post 1"
	post1.Slug = slug.Make(post1.Title)
	post2.Title = "A simple test title for post 2"
	post2.Slug = slug.Make(post2.Title)
	post3.Title = "A simple test title for post 3"
	post3.Slug = slug.Make(post3.Title)

	database.Instance.Create(&page)
	database.Instance.Create(&page1)
	database.Instance.Create(&page2)
	database.Instance.Create(&page3)

	database.Instance.Create(&post)
	database.Instance.Create(&post1)
	database.Instance.Create(&post2)
	database.Instance.Create(&post3)

	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/home", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, uint(1), page.ID)
	assert.Contains(t, w.Body.String(), "<title>Home</title>")
	assert.Contains(t, w.Body.String(), "<a href=\"/home\" class=\"logo\">Segantini's Blog</a>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home\" class=\"nav-link active\">Home</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/links\" class=\"nav-link\">Links</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/fale-conosco\" class=\"nav-link\">Fale conosco</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/sobre-nos\" class=\"nav-link\">Sobre nós</a></li>")
	assert.Contains(t, w.Body.String(), page.Body)
	assert.Contains(t, w.Body.String(), "<div>\n    Page 1/1\n</div>")
	assert.Contains(t, w.Body.String(), "<li class=\"page-item current\"><strong>1</strong></li>\n")
	assert.Contains(t, w.Body.String(), "<h5 class=\"card-title\">A simple test title for first post</h5>\n")
	assert.Contains(t, w.Body.String(), "<h5 class=\"card-title\">A simple test title for post 1</h5>\n")
	assert.Contains(t, w.Body.String(), "<h5 class=\"card-title\">A simple test title for post 2</h5>\n")
	assert.Contains(t, w.Body.String(), "<h5 class=\"card-title\">A simple test title for post 3</h5>\n")
	assert.Contains(t, w.Body.String(), "<small class=\"text-muted\">"+post.CreatedAt.Format("Jan 02, 2006")+"</small>")
	assert.Contains(t, w.Body.String(), "<a href=\"/home/articles/a-simple-test-title-for-first-post\" class=\"btn btn-primary\">Ler mais >></a>\n")
	assert.Contains(t, w.Body.String(), "<a href=\"/home/articles/a-simple-test-title-for-post-1\" class=\"btn btn-primary\">Ler mais >></a>\n")
	assert.Contains(t, w.Body.String(), "<a href=\"/home/articles/a-simple-test-title-for-post-2\" class=\"btn btn-primary\">Ler mais >></a>\n")
	assert.Contains(t, w.Body.String(), "<a href=\"/home/articles/a-simple-test-title-for-post-3\" class=\"btn btn-primary\">Ler mais >></a>\n")
}

func TestViewRouteFirst(t *testing.T) {
	//	database.Connect("root:root@tcp(localhost:3306)/cms-database?parseTime=true")
	database.Connect("")

	database.Migrate()

	var (
		page, page1, page2, page3 entity.PageMysql
	)

	page.Title = "Home"
	page.Slug = slug.Make(page.Title)
	page1.Title = "Links"
	page1.Slug = slug.Make(page1.Title)
	page1.Body = "<p>&nbsp;</p><figure class=\"image\"><img src=\"http://localhost:8000/uploads/undraw_writer_q06d.png\"></figure><p><strong>Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis.</strong></p>"
	page2.Title = "Sobre nós"
	page2.Slug = slug.Make(page2.Title)
	page2.Body = "a simple text body"
	page3.Title = "Fale conosco"
	page3.Slug = slug.Make(page3.Title)
	page3.Body = "a simple test text body"

	database.Instance.Create(&page)
	database.Instance.Create(&page1)
	database.Instance.Create(&page2)
	database.Instance.Create(&page3)

	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/home/links", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "<title>Links</title>")
	assert.Contains(t, w.Body.String(), page1.Body)
	assert.NotContains(t, w.Body.String(), "<title>Sobre nós</title>")
	assert.NotContains(t, w.Body.String(), page2.Body)
	assert.NotContains(t, w.Body.String(), "<title>Fale conosco</title>")
	assert.NotContains(t, w.Body.String(), page3.Body)

	assert.Contains(t, w.Body.String(), "<a href=\"/home\" class=\"logo\">Segantini's Blog</a>")
	//	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home\" class=\"nav-link\">Home</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/links\" class=\"nav-link active\">Links</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/fale-conosco\" class=\"nav-link\">Fale conosco</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/sobre-nos\" class=\"nav-link\">Sobre nós</a></li>")

}

func TestViewRouteSecond(t *testing.T) {
	//	database.Connect("root:root@tcp(localhost:3306)/cms-database?parseTime=true")
	database.Connect("")

	database.Migrate()

	var (
		page, page1, page2, page3 entity.PageMysql
	)

	page.Title = "Home"
	page.Slug = slug.Make(page.Title)
	page1.Title = "Links"
	page1.Slug = slug.Make(page1.Title)
	page1.Body = "<p>&nbsp;</p><figure class=\"image\"><img src=\"http://localhost:8000/uploads/undraw_writer_q06d.png\"></figure><p><strong>Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis.</strong></p>"
	page2.Title = "Sobre nós"
	page2.Slug = slug.Make(page2.Title)
	page2.Body = "a simple text body"
	page3.Title = "Fale conosco"
	page3.Slug = slug.Make(page3.Title)
	page3.Body = "a simple test text body"

	database.Instance.Create(&page)
	database.Instance.Create(&page1)
	database.Instance.Create(&page2)
	database.Instance.Create(&page3)

	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/home/"+page2.Slug, nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "<title>Sobre nós</title>")
	assert.Contains(t, w.Body.String(), page2.Body)
	assert.NotContains(t, w.Body.String(), "<title>Links</title>")
	assert.NotContains(t, w.Body.String(), page1.Body)
	assert.NotContains(t, w.Body.String(), "<title>Fale conosco</title>")
	assert.NotContains(t, w.Body.String(), page3.Body)

	assert.Contains(t, w.Body.String(), "<a href=\"/home\" class=\"logo\">Segantini's Blog</a>")
	//	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home\" class=\"nav-link\">Home</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/links\" class=\"nav-link\">Links</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/fale-conosco\" class=\"nav-link\">Fale conosco</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/sobre-nos\" class=\"nav-link active\">Sobre nós</a></li>")

}

func TestShowRoute(t *testing.T) {
	//	database.Connect("root:root@tcp(localhost:3306)/cms-database?parseTime=true")
	database.Connect("")

	database.Migrate()

	var (
		page, page1, page2, page3 entity.PageMysql
		post                      entity.PostMysql
	)

	page.Title = "Home"
	page.Slug = slug.Make(page.Title)
	page1.Title = "Links"
	page1.Slug = slug.Make(page1.Title)
	page1.Body = "<p>&nbsp;</p><figure class=\"image\"><img src=\"http://localhost:8000/uploads/undraw_writer_q06d.png\"></figure><p><strong>Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis.</strong></p>"
	page2.Title = "Sobre nós"
	page2.Slug = slug.Make(page2.Title)
	page2.Body = "a simple text body"
	page3.Title = "Fale conosco"
	page3.Slug = slug.Make(page3.Title)
	page3.Body = "a simple test text body"

	database.Instance.Create(&page)
	database.Instance.Create(&page1)
	database.Instance.Create(&page2)
	database.Instance.Create(&page3)

	post.Title = "A simple test title for first post"
	post.Slug = slug.Make(post.Title)
	post.Body = "<div><strong>A simple post title</strong> <p> A simple post introduction </p></div>"

	database.Instance.Create(&post)

	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/home/articles/"+post.Slug, nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), post.Body)
	assert.Contains(t, w.Body.String(), "<title>"+post.Title+"</title>")
	assert.NotContains(t, w.Body.String(), "<title>Sobre nós</title>")
	assert.NotContains(t, w.Body.String(), "<title>Links</title>")
	assert.NotContains(t, w.Body.String(), page1.Body)
	assert.NotContains(t, w.Body.String(), "<title>Fale conosco</title>")
	assert.NotContains(t, w.Body.String(), page3.Body)

	assert.Contains(t, w.Body.String(), "<a href=\"/home\" class=\"logo\">Segantini's Blog</a>")
	//assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home\" class=\"nav-link\">Home</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/links\" class=\"nav-link\">Links</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/fale-conosco\" class=\"nav-link\">Fale conosco</a></li>")
	assert.Contains(t, w.Body.String(), "<li class=\"nav-item\"><a href=\"/home/sobre-nos\" class=\"nav-link\">Sobre nós</a></li>")

}
