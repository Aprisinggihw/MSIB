package models

import (
	"fmt"
	"weekly-task-1/config"
	"weekly-task-1/entities"
)

// Fungsi GetAllProduct mengambil semua produk dari database
func GetAllProduct() ([]entities.Product, error) {
	query := `SELECT id, nama, kategori, stok, deskripsi FROM products;` // Query untuk mendapatkan semua produk
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal untuk get products: %w", err)
	}
	defer rows.Close() // Menutup rows setelah selesai

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Nama,
			&product.Kategori,
			&product.Stok,
			&product.Deskripsi)
		if err != nil {
			return nil, fmt.Errorf("gagal untuk scan product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// Fungsi GetProductById mengambil produk berdasarkan ID
func GetProductById(id int) (*entities.Product, error) {
	query := `SELECT id, nama, kategori, stok, deskripsi FROM products WHERE id = ?;` // Query untuk mendapatkan produk berdasarkan id
	rows, err := config.DB.Query(query, id)
	if err != nil {

		return nil, fmt.Errorf("gagal untuk get products: %w", err)
	}
	defer rows.Close() // Menutup rows setelah selesai
	product := &entities.Product{}

	if rows.Next() {
		// Scan hasil query ke dalam product
		err = rows.Scan(
			&product.Id,
			&product.Nama,
			&product.Kategori,
			&product.Stok,
			&product.Deskripsi,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal untuk scan product: %w", err)
		}
	} else {
		return nil, fmt.Errorf("produk dengan id %d tidak ditemukan", id)
	}

	return product, nil
}

// Menambahkan produk baru ke dalam database
func AddDataProduct(product *entities.Product) error {
	query := `INSERT INTO products (nama, kategori, stok, deskripsi) VALUES (?, ?, ?, ?);` // Query untuk menambahkan produk
	_, err := config.DB.Exec(query, product.Nama, product.Kategori, product.Stok, product.Deskripsi)
	if err != nil {
		return fmt.Errorf("gagal untuk insert product: %w", err)
	}
	return nil
}

// Fungsi EditDataProduct memperbarui produk yang sudah ada dalam database
func EditDataProduct(id int, product entities.Product) error {
	query := `UPDATE products SET 
		nama = ?, 
		kategori = ?, 
		stok = ?, 
		deskripsi = ?  
		WHERE id = ?;` // Query untuk memperbarui produk berdasarkan id

	_, err := config.DB.Exec(query, product.Nama, product.Kategori, product.Stok, product.Deskripsi, id)
	if err != nil {
		return fmt.Errorf("gagal untuk update product: %w", err)
	}

	return nil
}

// Delete menghapus produk dari database berdasarkan id
func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM products WHERE id = ?`, id) // Menjalankan query untuk menghapus produk
	if err != nil {
		return fmt.Errorf("gagal untuk delete product: %w", err)
	}
	return nil
}
