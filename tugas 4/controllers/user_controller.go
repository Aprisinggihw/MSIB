package controller

import (
	"net/http"
	"strconv"
	"tugas-4/entities"
	"tugas-4/models"

	"github.com/labstack/echo/v4"
)

func CreateUser(ctx echo.Context) error {
	user := new(entities.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "infailed to bind",
			"data": nil,
		})
	}

	err := models.CreateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "create user failed",
			"data": nil,
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "sukses created product",
		"data":    user,
	})
}
// GetAllUsers - Endpoint untuk menampilkan semua user
func GetAllUsers(ctx echo.Context) error {
	users, err := models.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to fetch users",
			"data": nil,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "successfully retrieved users",
			"data": users,
		})
}


// UpdateUser - Endpoint untuk memperbarui user berdasarkan ID
func UpdateUser(ctx echo.Context) error {
	// Mendapatkan ID dari parameter URL
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user id",
			"data":    nil,
		})}

	// Mencari user berdasarkan ID
	user, err := models.FindUserByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"data":    nil,
		})
	}

	// Mengambil data user baru dari request body
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "infailed to bind",
			"data": nil,
		})
	}

	// Memperbarui user
	err = models.UpdateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update user",
			"data": nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated successfully",
		"data":    user,
	})
}

func DeleteUser(ctx echo.Context) error {
	// Mendapatkan ID dari parameter URL
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user id",
			"data":    nil,
		})
	}

	// Mencari user berdasarkan ID
	user, err := models.FindUserByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"data":    nil,
		})
	}

	// Menghapus user
	err = models.DeleteUser(user.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete user",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "user deleted successfully",
		"data":    user,
	})
}
