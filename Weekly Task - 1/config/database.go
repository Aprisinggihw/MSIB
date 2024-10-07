package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DB adalah variabel global yang menyimpan koneksi ke database
var DB *sql.DB

// Fungsi ConnectDB membuat koneksi ke database SQLite dan menyimpannya dalam variabel global DB
func ConnectDB() error {
	// Membuka koneksi ke database SQLite
	db, err := sql.Open("sqlite3", "C:/sqlite3/myDatabase/golang_product.db")
	if err != nil {

		return fmt.Errorf("gagal untuk koneksi ke database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close() // Menutup koneksi jika ping gagal
		return fmt.Errorf("gagal untuk ping database: %w", err)
	}

	// Menyimpan koneksi yang berhasil ke variabel global DB
	DB = db
	return nil
}
