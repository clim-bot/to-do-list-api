package main

import (

    "github.com/gin-gonic/gin"

    "github.com/clim-bot/todo-list-api/config"
    "github.com/clim-bot/todo-list-api/models"
    "github.com/clim-bot/todo-list-api/routes"
)


func main() {
    db := config.InitDB()
    db.AutoMigrate(&models.User{}, &models.Task{})

    router := gin.Default()
    routes.SetupRoutes(router, db)
    router.Run(":8080")
}
