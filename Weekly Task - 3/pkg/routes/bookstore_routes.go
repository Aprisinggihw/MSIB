package routes

import (
	"weekly-task-3/pkg/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterBookRoutes(e *echo.Echo, db *gorm.DB) {
	bookController := controllers.BookController{DB: db}

	e.GET("/book", bookController.GetBooks)
	e.POST("/book", bookController.CreateBook)
	e.GET("/book/:bookId", bookController.GetBookByID)
	e.PUT("/book/:bookId", bookController.UpdateBook)
	e.DELETE("/book/:bookId", bookController.DeleteBook)
}
