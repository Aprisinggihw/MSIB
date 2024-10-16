package config

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// Format DSN untuk MySQL
	dsn := "root:singgih1@tcp(127.0.0.1:3306)/tugas4?charset=utf8mb4&parseTime=True&loc=Local"
	
	// Koneksi ke MySQL menggunakan GORM
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
}

