package controller

import (
	"net/http"
	"strconv"
	"tugas-4/entities"
	"tugas-4/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateTodo(ctx echo.Context) error {
	todo := new(entities.Todo)
	if err := ctx.Bind(todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind",
			"data": nil,
		})
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	todo.UserID = uint(claims["user_id"].(float64))

	err := models.CreateTodo(todo)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "create todo failed",
			"data":nil,
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "sukses created product",
		"data":    todo,
	})
}

func GetTodos(ctx echo.Context) error {
	todos, err := models.GetTodos()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "get todos failed",
			"data": nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get todos",
		"data":    todos,
	})
}

// UpdateTodo - Endpoint untuk memperbarui todo berdasarkan ID
func UpdateTodo(ctx echo.Context) error {
	// Mendapatkan ID dari parameter URL
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid todo id",
			"data":    nil,
		})
	}

	// Mencari todo berdasarkan ID
	todo, err := models.FindTodoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "todo not found",
			"data":    nil,
		})
	}

	// Mengambil data todo baru dari request body
	if err := ctx.Bind(todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "infailed to bind",
			"data": nil,
		})
	}

	// Memperbarui todo
	err = models.UpdateTodo(todo)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update todo",
			"data": nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated successfully",
		"data":    todo,
	})
}


// DeleteTodo - Endpoint untuk menghapus todo berdasarkan ID
func DeleteTodo(ctx echo.Context) error {
	// Mendapatkan ID dari parameter URL
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid todo id",
			"data":    nil,
		})
	}

	// Mencari todo berdasarkan ID
	todo, err := models.FindTodoByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "todo not found",
			"data":    nil,
		})
	}

	// Menghapus todo
	err = models.DeleteTodo(todo.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete todo",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "todo deleted successfully",
		"data":    todo,
	})
}

