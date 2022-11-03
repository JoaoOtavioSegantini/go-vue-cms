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
	"github.com/joaotavioos/cms-server/auth"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
	"github.com/joaotavioos/cms-server/routes"
	"github.com/stretchr/testify/assert"
)

func SetupTestAuth() (string, error) {
	var user entity.UserMysql

	user.Username = "john_doe"
	user.Password = "123"
	user.Email = "john_doe@gmail.com"
	user.Name = "John Doe"

	user.HashPassword(user.Password)

	database.Instance.Create(&user)

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func TestPageUnauthenticated(t *testing.T) {
	t.Parallel()
	database.Connect("")

	database.Migrate()
	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/res-data", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	//	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, w.Body.String(), "{\"error\":\"request does not contain an access token\"}")

}

func TestPageWhenPageIsEmpty(t *testing.T) {
	database.Connect("")

	database.Migrate()

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/res-data", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	//	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), "[]")

}

func TestFindAllRequest(t *testing.T) {
	database.Connect("")

	database.Migrate()

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	var (
		page, page1, page2, page3 entity.PageMysql
	)

	page.Title = "Home"
	page.Slug = slug.Make(page.Title)
	page1.Title = "Links"
	page1.Slug = slug.Make(page1.Title)
	page1.Body = "Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis."
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

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/res-data", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)

	r.ServeHTTP(w, req)

	mockResponse := fmt.Sprintf(
		`[{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"title":"%s","body":"%s","slug":"%s"},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"title":"%s","body":"%s","slug":"%s"},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"title":"%s","body":"%s","slug":"%s"},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"title":"%s","body":"%s","slug":"%s"}]`,
		page.ID, page.CreatedAt.Format(time.RFC3339Nano), page.UpdatedAt.Format(time.RFC3339Nano), page.Title, page.Body, page.Slug,
		page1.ID, page1.CreatedAt.Format(time.RFC3339Nano), page1.UpdatedAt.Format(time.RFC3339Nano), page1.Title, page1.Body, page1.Slug,
		page2.ID, page2.CreatedAt.Format(time.RFC3339Nano), page2.UpdatedAt.Format(time.RFC3339Nano), page2.Title, page2.Body, page2.Slug,
		page3.ID, page3.CreatedAt.Format(time.RFC3339Nano), page3.UpdatedAt.Format(time.RFC3339Nano), page3.Title, page3.Body, page3.Slug,
	)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), mockResponse)

}

func TestSave(t *testing.T) {
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

	page := entity.PageMysql{
		Title: "page 1",
		Body:  "Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis.",
		Slug:  slug.Make("page 1"),
	}
	jsonValue, _ := json.Marshal(page)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/v1/res-data", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), page.Body)
	assert.Contains(t, w.Body.String(), page.Title)
	assert.Contains(t, w.Body.String(), page.Slug)

}

func TestPageUpdate(t *testing.T) {
	t.Parallel()
	database.Connect("")

	database.Migrate()

	var page entity.PageMysql

	page.Title = "Home"
	page.Slug = slug.Make(page.Title)

	database.Instance.Create(&page)

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	page = entity.PageMysql{
		Title: "page 1 updated",
		Body:  "Um blog updated",
		Slug:  slug.Make("page 1 updated"),
	}
	jsonValue, _ := json.Marshal(page)

	req, err := http.NewRequestWithContext(ctx, "PUT", "/api/v1/res-data/1", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), page.Body)
	assert.Contains(t, w.Body.String(), page.Title)
	assert.Contains(t, w.Body.String(), page.Slug)

}

func TestPageFindById(t *testing.T) {
	database.Connect("")

	database.Migrate()

	var page entity.PageMysql

	page.Title = "Home"
	page.Slug = slug.Make(page.Title)
	page.Body = "Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis."

	database.Instance.Create(&page)

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/res-data/1", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), page.Body)
	assert.Contains(t, w.Body.String(), page.Title)
	assert.Contains(t, w.Body.String(), page.Slug)

}

func TestPageDelete(t *testing.T) {
	database.Connect("")

	database.Migrate()

	var page entity.PageMysql
	var count int64

	page.Title = "Home"
	page.Slug = slug.Make(page.Title)
	page.Body = "Um blog sobre assuntos diversos, tais como tecnologia, ciência e leis."

	database.Instance.Create(&page)
	database.Instance.Model(&page).Count(&count)

	assert.Equal(t, count, int64(1))

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "DELETE", "/api/v1/res-data/1", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	database.Instance.Model(&page).Count(&count)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, count, int64(0))

}
