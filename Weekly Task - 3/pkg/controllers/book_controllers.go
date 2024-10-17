package controllers

import (
	"net/http"

	"weekly-task-3/pkg/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

// Membuat Buku Baru
func (bc *BookController) CreateBook(ctx echo.Context) error {
	var book models.Book
	if err := ctx.Bind(&book); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid input",
			"data": nil,
		})
	}

	if err := bc.DB.Create(&book).Error; err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "failed to create",
            "data": nil,
        })
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
        "message": "successfully created",
        "data": book,
    })
}

// Mendapatkan List Buku
func (bc *BookController) GetBooks(ctx echo.Context) error {
	var books []models.Book
	if err := bc.DB.Find(&books).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "failed to find books",
            "data": nil,
        })
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
        "message": "successfully find books",
        "data": books,
    })
}

// Mendapatkan Buku Berdasarkan ID
func (bc *BookController) GetBookByID(ctx echo.Context) error {
	id := ctx.Param("bookId")
	var book models.Book
	if err := bc.DB.First(&book, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
            "message": "book not found",
            "data": nil,
        })
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
        "message": "successfully found book",
        "data": book,
    })
}

// Memperbarui Buku Berdasarkan ID
func (bc *BookController) UpdateBook(ctx echo.Context) error {
	id := ctx.Param("bookId")
	var book models.Book
	if err := bc.DB.First(&book, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
            "message": "book not found",
            "data": nil,
        })
	}
    var newBook models.Book

	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
            "message": "invalid input",
            "data": nil,
        })
	}

    // Update hanya field yang diubah dari data baru
	if newBook.Title != "" {
		book.Title = newBook.Title
	}
	if newBook.Author != "" {
		book.Author = newBook.Author
	}
	if newBook.Stock != 0 {
		book.Stock = newBook.Stock
	}
	if newBook.Price != 0 {
		book.Price = newBook.Price
	}

	if err := bc.DB.Save(&book).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "failed to save book",
            "data": nil,
        })
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
        "message": "successfully saved book",
        "data": book,
    })
}

// Menghapus Buku Berdasarkan ID
func (bc *BookController) DeleteBook(ctx echo.Context) error {
	id := ctx.Param("bookId")
    var book models.Book
    if err := bc.DB.First(&book, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
            "message": "book not found",
            "data": nil,
        })
	}
	if err := bc.DB.Delete(&models.Book{}, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
            "message": "book not found",
            "data": nil,
        })
	}
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "message": "success delete book",
        "data": book,
    })
}
