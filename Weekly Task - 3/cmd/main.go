package main

import (
    "weekly-task-3/pkg/config"
    "weekly-task-3/pkg/models"
    "weekly-task-3/pkg/routes"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // Inisialisasi Database
    db := config.InitDB()
    db.AutoMigrate(&models.Book{})

    // Register Routes
    routes.RegisterBookRoutes(e, db)

    // Start Server
    e.Logger.Fatal(e.Start(":8080"))
}
