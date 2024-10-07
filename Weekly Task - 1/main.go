package main

import (
	"fmt"
	"log"
	"weekly-task-1/config"
	controller "weekly-task-1/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Menghubungkan ke database
	err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	defer config.DB.Close() // Menutup koneksi database saat program selesai

	e := echo.New()

	e.GET("/", controller.GettAllProduct)             // Untuk mengambil semua produk
	e.POST("/add", controller.AddProduct)             // Untuk menambahkan produk baru
	e.PUT("/edit/:id", controller.EditProduct)        // Untuk memperbarui produk berdasarkan ID
	e.DELETE("/delete/:id", controller.DeleteProduct) // Untuk menghapus produk berdasarkan ID

	log.Println("Server berjalan pada port 8080")

	e.Logger.Fatal(e.Start(":8080"))
}
