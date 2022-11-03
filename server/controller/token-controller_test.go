package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joaotavioos/cms-server/controller"
	"github.com/joaotavioos/cms-server/database"
	"github.com/joaotavioos/cms-server/entity"
	"github.com/joaotavioos/cms-server/routes"
	"github.com/stretchr/testify/assert"
)

func TestUserAuthSuccess(t *testing.T) {
	var user entity.UserMysql

	database.Connect("")

	database.Migrate()

	user.Username = "john_doe"
	user.Password = "123"
	user.Email = "john_doe@gmail.com"
	user.Name = "John Doe"

	user.HashPassword(user.Password)

	database.Instance.Create(&user)

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	request := controller.TokenRequest{
		Email:    user.Email,
		Password: "123",
	}

	jsonValue, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/token", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), "{\"token\":")

}

func TestUserAuthFailWhenEmailNotExists(t *testing.T) {
	var user entity.UserMysql

	database.Connect("")

	database.Migrate()

	user.Username = "john_doe"
	user.Password = "123"
	user.Email = "john_doe@gmail.com"
	user.Name = "John Doe"

	user.HashPassword(user.Password)

	database.Instance.Create(&user)

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	request := controller.TokenRequest{
		Email:    "fake_email@gmail.com",
		Password: "123",
	}

	jsonValue, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/token", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Exactly(t, w.Body.String(), "{\"error\":\"record not found\"}")

}

func TestUserAuthFailWhenPasswordNotExists(t *testing.T) {
	var user entity.UserMysql

	database.Connect("")

	database.Migrate()

	user.Username = "john_doe"
	user.Password = "123"
	user.Email = "john_doe@gmail.com"
	user.Name = "John Doe"

	user.HashPassword(user.Password)

	database.Instance.Create(&user)

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	routes.SetupRouter(r, "../templates/*.html")

	request := controller.TokenRequest{
		Email:    user.Email,
		Password: "fake password",
	}

	jsonValue, _ := json.Marshal(request)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/token", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Exactly(t, w.Body.String(), "{\"error\":\"invalid credentials\"}")

}
