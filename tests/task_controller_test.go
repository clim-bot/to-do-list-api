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

func TestCreateTask(t *testing.T) {
    db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    db.AutoMigrate(&models.Task{})

    taskController := &controllers.TaskController{DB: db}
    router := gin.Default()
    router.POST("/tasks", taskController.CreateTask)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/tasks", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTasks(t *testing.T) {
    db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    db.AutoMigrate(&models.Task{})

    taskController := &controllers.TaskController{DB: db}
    router := gin.Default()
    router.GET("/tasks", taskController.GetTasks)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/tasks", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}
