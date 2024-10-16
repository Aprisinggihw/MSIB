package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware adalah middleware untuk validasi JWT
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if strings.HasPrefix(tokenString, "Bearer ") {
				tokenString = tokenString[len("Bearer "):]
			}

			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing or malformed JWT"})
			}

			// Parsing token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Verifikasi metode signing
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.ErrUnauthorized
				}
				return []byte("secret"), nil 
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid JWT"})
			}

			// Simpan klaim ke dalam context
			c.Set("user", token)

			// Jika valid, lanjutkan ke handler berikutnya
			return next(c)
		}
	}
}

// IsAdmin adalah middleware untuk memastikan pengguna adalah admin
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims, ok := user.Claims.(jwt.MapClaims)
		if !ok || claims["role"] != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

// IsEditor adalah middleware untuk memastikan pengguna adalah editor
func IsEditor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims, ok := user.Claims.(jwt.MapClaims)
		if !ok || claims["role"] != "editor" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
