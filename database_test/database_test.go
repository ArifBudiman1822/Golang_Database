package database_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost;3306)/golang_database")
	if err != nil {
		panic(err)
	}

	 
	// gunakan db
	defer db.Close()
}