package main

import (
	"tugas-4/config"
	"tugas-4/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// Inisialisasi koneksi database menggunakan GORM
	config.ConnectDB()

	// Inisialisasi Echo
	e := echo.New()

	// Setup Routes
	routes.InitRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
