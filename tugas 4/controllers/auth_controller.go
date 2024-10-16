package controller

import (
	"net/http"
	"time"
	"tugas-4/entities"
	"tugas-4/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(ctx echo.Context) error {
	var req entities.LoginRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	username := req.Username
	password := req.Password

	user, err := models.FindUserByUsernameAndPassword(username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username and password are incorrect",
			"data":    nil,
		})
	}
	if password != user.Password {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username and password are incorrect",
			"data":    nil,
		})
	}

	// Generate JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "login successful",
		"data":    user,
		"token":   t,
	})
}
