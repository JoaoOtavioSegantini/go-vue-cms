package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
	"github.com/joaotavioos/cms-server/routes"
	"github.com/stretchr/testify/assert"
)

func TestPostUnauthenticated(t *testing.T) {
	t.Parallel()
	database.Connect("")

	database.Migrate()
	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/site-admin-posts", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	//	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, w.Body.String(), "{\"error\":\"request does not contain an access token\"}")

}

func TestPostWhenPageIsEmpty(t *testing.T) {
	database.Connect("")

	database.Migrate()

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/site-admin-posts", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	//	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), "[]")

}

func TestPostFindAllRequest(t *testing.T) {
	database.Connect("")

	database.Migrate()

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	var (
		post, post1, post2, post3 entity.PostMysql
	)

	post.Title = "A simple post title"
	post.Slug = slug.Make(post.Title)
	post1.Title = "A post title 1"
	post1.Slug = slug.Make(post1.Title)
	post1.Body = "Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis."
	post2.Title = "A post title 2"
	post2.Slug = slug.Make(post2.Title)
	post2.Body = "a simple text body"
	post3.Title = "Fale conosco"
	post3.Slug = slug.Make(post3.Title)
	post3.Body = "a simple test text body"

	database.Instance.Create(&post)
	database.Instance.Create(&post1)
	database.Instance.Create(&post2)
	database.Instance.Create(&post3)

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/site-admin-posts", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)

	r.ServeHTTP(w, req)

	mockResponse := fmt.Sprintf(
		`[{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"title":"%s","body":"%s","slug":"%s"},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"title":"%s","body":"%s","slug":"%s"},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"title":"%s","body":"%s","slug":"%s"},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"title":"%s","body":"%s","slug":"%s"}]`,
		post.ID, post.CreatedAt.Format(time.RFC3339Nano), post.UpdatedAt.Format(time.RFC3339Nano), post.Title, post.Body, post.Slug,
		post1.ID, post1.CreatedAt.Format(time.RFC3339Nano), post1.UpdatedAt.Format(time.RFC3339Nano), post1.Title, post1.Body, post1.Slug,
		post2.ID, post2.CreatedAt.Format(time.RFC3339Nano), post2.UpdatedAt.Format(time.RFC3339Nano), post2.Title, post2.Body, post2.Slug,
		post3.ID, post3.CreatedAt.Format(time.RFC3339Nano), post3.UpdatedAt.Format(time.RFC3339Nano), post3.Title, post3.Body, post3.Slug,
	)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), mockResponse)

}

func TestPostSave(t *testing.T) {
	t.Parallel()
	database.Connect("")

	database.Migrate()

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	post := entity.PostMysql{
		Title: "post 1",
		Body:  "Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis.",
		Slug:  slug.Make("post 1"),
	}
	jsonValue, _ := json.Marshal(post)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/v1/res-data", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), post.Body)
	assert.Contains(t, w.Body.String(), post.Title)
	assert.Contains(t, w.Body.String(), post.Slug)

}

func TestPostUpdate(t *testing.T) {
	database.Connect("")

	database.Migrate()

	var post entity.PostMysql

	post.Title = "A simple post title"
	post.Slug = slug.Make(post.Title)

	database.Instance.Create(&post)

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	post = entity.PostMysql{
		Title: "post 1 updated",
		Body:  "Um blog post updated",
		Slug:  slug.Make("post 1 updated"),
	}
	jsonValue, _ := json.Marshal(post)

	req, err := http.NewRequestWithContext(ctx, "PUT", "/api/v1/site-admin-posts/1", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), post.Body)
	assert.Contains(t, w.Body.String(), post.Title)
	assert.Contains(t, w.Body.String(), post.Slug)

}

func TestPostFindById(t *testing.T) {
	database.Connect("")

	database.Migrate()

	var post entity.PostMysql

	post.Title = "A simple post title"
	post.Slug = slug.Make(post.Title)
	post.Body = "Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis."

	database.Instance.Create(&post)

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/site-admin-posts/1", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), post.Body)
	assert.Contains(t, w.Body.String(), post.Title)
	assert.Contains(t, w.Body.String(), post.Slug)

}

func TestPostDelete(t *testing.T) {
	database.Connect("")

	database.Migrate()

	var post entity.PostMysql
	var count int64

	post.Title = "A simple post title"
	post.Slug = slug.Make(post.Title)
	post.Body = "Um post de um blog sobre assuntos diversos, tais como tecnologia, ciência e leis."

	database.Instance.Create(&post)
	database.Instance.Model(&post).Count(&count)

	assert.Equal(t, count, int64(1))

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "DELETE", "/api/v1/site-admin-posts/1", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	database.Instance.Model(&post).Count(&count)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, count, int64(0))

}
