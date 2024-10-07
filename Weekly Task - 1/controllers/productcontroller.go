package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"weekly-task-1/entities"
	"weekly-task-1/models"

	"github.com/labstack/echo/v4"
)

// GettAllProduct mengambil semua produk dari database dan mengembalikannya dalam format JSON
func GettAllProduct(ctx echo.Context) error {

	products, err := models.GetAllProduct()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": fmt.Sprint(err),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses mengambil semua data product",
		"data":    products,
	})
}

// AddProduct menambahkan produk baru ke database berdasarkan data dari request
func AddProduct(ctx echo.Context) error {
	productRequest := new(entities.Product)

	if err := ctx.Bind(productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "gagal untuk bind request data",
		})
	}
	// Memanggil model untuk menambahkan produk
	err := models.AddDataProduct(productRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": fmt.Sprint(err),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses created product",
	})
}

// EditProduct memperbarui informasi produk berdasarkan id yang diterima dari URL
func EditProduct(ctx echo.Context) error {
	// Mengambil parameter id dari URL dan mengonversinya ke integer
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "product id tidak valid",
		})
	}

	// Mengambil data produk lama dari database berdasarkan id
	dataProduct, err := models.GetProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprint(err),
		})
	}

	var dataBaru entities.Product

	if err := ctx.Bind(&dataBaru); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "gagal untuk bind request data",
		})
	}

	// Update hanya field yang diubah dari data baru
	if dataBaru.Nama != "" {
		dataProduct.Nama = dataBaru.Nama
	}
	if dataBaru.Kategori != "" {
		dataProduct.Kategori = dataBaru.Kategori
	}
	if dataBaru.Stok != 0 {
		dataProduct.Stok = dataBaru.Stok
	}
	if dataBaru.Deskripsi != "" {
		dataProduct.Deskripsi = dataBaru.Deskripsi
	}

	// Memanggil fungsi model untuk update data produk
	if err := models.EditDataProduct(id, *dataProduct); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": fmt.Sprint(err),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "sukses memperbarui data product",
	})
}

// DeleteProduct menghapus produk berdasarkan id yang diterima dari URL
func DeleteProduct(ctx echo.Context) error {
	// Mengambil parameter id dari URL dan mengonversinya ke integer
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "product id tidak valid",
		})
	}
	// Memanggil model untuk menghapus produk
	if err := models.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": fmt.Sprint(err),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "sukses menghapus data product",
	})
}
