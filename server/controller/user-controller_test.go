package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
	"github.com/joaotavioos/cms-server/routes"
	"github.com/stretchr/testify/assert"
)

func TestUserUnauthenticated(t *testing.T) {
	database.Connect("")

	database.Migrate()
	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/users-site-admin", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	//	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, w.Body.String(), "{\"error\":\"request does not contain an access token\"}")

}

func TestUserWhenIsNotEmpty(t *testing.T) {
	database.Connect("")

	database.Migrate()

	database.Instance.Where("1 = 1").Delete(&entity.UserMysql{})

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/users-site-admin", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	//	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())

}

func TestUserFindAllRequest(t *testing.T) {
	database.Connect("")

	database.Migrate()

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	var (
		user, user1, user2, user3, authUser entity.UserMysql
	)

	database.Instance.Where("ID = ?", "1").Find(&authUser)

	user.Name = "A simple username title"
	user.Email = "user@user.com"
	user1.Name = "A user title 1"
	user1.Email = "user@user1.com"
	user1.Username = "Um simple username"
	user2.Name = "A simple username 2 title"
	user2.Email = "user@user2.com"
	user2.Username = "a simple username text"
	user3.Name = "A simple username 3 title"
	user3.Email = "user@user3.com"
	user3.Username = "a simple user test text username"

	database.Instance.Create(&user)
	database.Instance.Create(&user1)
	database.Instance.Create(&user2)
	database.Instance.Create(&user3)

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/users-site-admin", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)

	r.ServeHTTP(w, req)

	mockResponse := fmt.Sprintf(
		`[{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"name":"%s","username":"%s","email":"%s","password":""},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"name":"%s","username":"%s","email":"%s","password":""},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"name":"%s","username":"%s","email":"%s","password":""},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"name":"%s","username":"%s","email":"%s","password":""},{"ID":%d,"CreatedAt":"%s","UpdatedAt":"%s","DeletedAt":null,"name":"%s","username":"%s","email":"%s","password":""}]`,
		authUser.ID, authUser.CreatedAt.Format(time.RFC3339Nano), authUser.UpdatedAt.Format(time.RFC3339Nano), authUser.Name, authUser.Username, authUser.Email,
		user.ID, user.CreatedAt.Format(time.RFC3339Nano), user.UpdatedAt.Format(time.RFC3339Nano), user.Name, user.Username, user.Email,
		user1.ID, user1.CreatedAt.Format(time.RFC3339Nano), user1.UpdatedAt.Format(time.RFC3339Nano), user1.Name, user1.Username, user1.Email,
		user2.ID, user2.CreatedAt.Format(time.RFC3339Nano), user2.UpdatedAt.Format(time.RFC3339Nano), user2.Name, user2.Username, user2.Email,
		user3.ID, user3.CreatedAt.Format(time.RFC3339Nano), user3.UpdatedAt.Format(time.RFC3339Nano), user3.Name, user3.Username, user3.Email,
	)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), mockResponse)

}

func TestUserSave(t *testing.T) {
	database.Connect("")

	database.Migrate()

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	user := entity.UserMysql{
		Name:     "user 1",
		Username: "user1_user",
		Email:    "user@user.com",
		Password: "123",
	}
	jsonValue, _ := json.Marshal(user)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/v1/users-site-admin", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), user.Username)
	assert.Contains(t, w.Body.String(), user.Email)
	//assert.Contains(t, w.Body.String(), 2)

}

func TestUserUpdate(t *testing.T) {
	database.Connect("")

	database.Migrate()

	var user entity.UserMysql

	user.Name = "A simple user title"
	user.Email = slug.Make(user.Name)

	database.Instance.Create(&user)

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	user = entity.UserMysql{
		Name:     "user 1",
		Username: "user1_user",
		Email:    "user@user.com",
		Password: "123",
	}
	jsonValue, _ := json.Marshal(user)

	req, err := http.NewRequestWithContext(ctx, "PUT", "/api/v1/users-site-admin/1", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), user.Name)
	assert.Contains(t, w.Body.String(), user.Email)
	assert.Contains(t, w.Body.String(), user.Username)

}

func TestUserFindById(t *testing.T) {
	database.Connect("")

	database.Migrate()

	user := entity.UserMysql{
		Name:     "user 1",
		Username: "user1_user",
		Email:    "user@user.com",
		Password: "123",
	}

	database.Instance.Create(&user)

	tokenString, err := SetupTestAuth()

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/users-site-admin/1", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), user.Name)
	assert.Contains(t, w.Body.String(), user.Email)
	assert.Contains(t, w.Body.String(), user.Username)

}

func TestUserDelete(t *testing.T) {
	database.Connect("")

	database.Migrate()

	var user entity.UserMysql
	var count int64

	user = entity.UserMysql{
		Name:     "user 1",
		Username: "user1_user",
		Email:    "user@user.com",
		Password: "123",
	}

	database.Instance.Create(&user)

	tokenString, err := SetupTestAuth()
	database.Instance.Model(&user).Count(&count)

	assert.Equal(t, count, int64(2))

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	req, err := http.NewRequestWithContext(ctx, "DELETE", "/api/v1/users-site-admin/1", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Authorization", tokenString)
	r.ServeHTTP(w, req)

	database.Instance.Model(&user).Count(&count)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, count, int64(1))

}

func TestRegisterUserUnauthenticated(t *testing.T) {
	database.Connect("")

	database.Migrate()
	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	user := entity.UserMysql{
		Name:     "fake",
		Email:    "fake@gmail.com",
		Password: "fake",
		Username: "User fake",
	}

	jsonValue, _ := json.Marshal(user)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/user/register", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	//	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, w.Body.String(), "")

}

func TestRegisterUserAuthenticate(t *testing.T) {
	database.Connect("")

	database.Migrate()
	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	user := entity.UserMysql{
		Name:     "fake",
		Email:    "fake@gmail.com",
		Password: "fake",
		Username: "User fake",
	}

	jsonValue, _ := json.Marshal(user)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/user/register", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("USER"), os.Getenv("PASSWORD"))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Exactly(t, w.Body.String(), "{\"email\":\"fake@gmail.com\",\"userId\":1,\"username\":\"User fake\"}")

}
