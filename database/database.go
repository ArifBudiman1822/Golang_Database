package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // Minimal koneksi yang terbuka
	db.SetMaxOpenConns(100)                 // Maximal open yang terbuka
	db.SetConnMaxIdleTime(5 * time.Minute)  // Brp lama koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // Brp lama koneksi dapat digunakan

	return db
}
