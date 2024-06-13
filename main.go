package main

import (

    "github.com/gin-gonic/gin"

    "github.com/clim-bot/todo-list-api/config"
    "github.com/clim-bot/todo-list-api/models"
    "github.com/clim-bot/todo-list-api/routes"
    "github.com/gin-contrib/cors"
)


func main() {
    db := config.InitDB()
    db.AutoMigrate(&models.User{}, &models.Task{})

    router := gin.Default()

    // CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    routes.SetupRoutes(router, db)
    router.Run(":8080")
}