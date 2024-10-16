package routes

import (
    "tugas-4/controllers" 
    "tugas-4/middleware"
    "github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
    // Login route
    e.POST("/login", controller.Login)

    // Group with JWT Middleware
    api := e.Group("/api")
    api.Use(middleware.JWTMiddleware())

    // Todo routes (Editor only)
    todos := api.Group("/todos")
    todos.Use(middleware.IsEditor) 
    todos.POST("", controller.CreateTodo)
    todos.GET("", controller.GetTodos)
    todos.PUT("/:id", controller.UpdateTodo)
    todos.DELETE("/:id", controller.DeleteTodo)

    // User routes (Admin only)
    users := api.Group("/users")
    users.Use(middleware.IsAdmin) 
    users.POST("", controller.CreateUser)
    users.GET("", controller.GetAllUsers)
    users.PUT("/:id", controller.UpdateUser)
    users.DELETE("/:id", controller.DeleteUser)
}
