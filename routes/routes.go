package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/clim-bot/todo-list-api/controllers"
    "github.com/clim-bot/todo-list-api/middleware"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    authController := &controllers.AuthController{DB: db}
    taskController := &controllers.TaskController{DB: db}

    authRoutes := router.Group("/auth")
    {
        authRoutes.POST("/register", authController.Register)
        authRoutes.POST("/login", authController.Login)
    }

    taskRoutes := router.Group("/tasks")
    taskRoutes.Use(middleware.AuthMiddleware())
    {
        taskRoutes.POST("", taskController.CreateTask)
        taskRoutes.GET("", taskController.GetTasks)
        taskRoutes.PUT("", taskController.UpdateTask)
        taskRoutes.DELETE("", taskController.DeleteTask)
    }
}
