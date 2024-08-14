package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import dialect mysql untuk gorm
)

// Fungsi ConnectDB mengembalikan koneksi database GORM yang baru
func ConnectDB() (*gorm.DB, error) {
	// Pengaturan koneksi database
	host := "localhost"
	port := "3306"
	username := "root"
	password := ""
	dbname := "tugas_1"

	// Buat string DSN (Data Source Name) untuk koneksi database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// Buka koneksi database baru menggunakan GORM
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		// Kembalikan error jika koneksi tidak dapat dibuat
		return nil, err
	}

	// Kembalikan koneksi database yang baru
	return db, nil
}