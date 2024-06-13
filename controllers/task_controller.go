package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/clim-bot/todo-list-api/models"
    "github.com/clim-bot/todo-list-api/repositories"
    "github.com/clim-bot/todo-list-api/utils"
)

type TaskController struct {
    DB *gorm.DB
}

func (tc *TaskController) CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userEmail, _ := utils.ExtractUserEmailFromJWT(c)

    user, err := repositories.GetUserByEmail(tc.DB, userEmail)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    task.UserID = user.ID

    if err := repositories.CreateTask(tc.DB, &task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
    userEmail, _ := utils.ExtractUserEmailFromJWT(c)

    user, err := repositories.GetUserByEmail(tc.DB, userEmail)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    tasks, err := repositories.GetTasksByUserID(tc.DB, user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := repositories.UpdateTask(tc.DB, &task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := repositories.DeleteTask(tc.DB, &task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
