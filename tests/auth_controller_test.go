package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/clim-bot/todo-list-api/controllers"
    "github.com/clim-bot/todo-list-api/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func TestRegister(t *testing.T) {
    db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    db.AutoMigrate(&models.User{})

    authController := &controllers.AuthController{DB: db}
    router := gin.Default()
    router.POST("/auth/register", authController.Register)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/auth/register", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin(t *testing.T) {
    db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    db.AutoMigrate(&models.User{})

    authController := &controllers.AuthController{DB: db}
    router := gin.Default()
    router.POST("/auth/login", authController.Login)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/auth/login", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}
