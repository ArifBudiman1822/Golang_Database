package sql_test

import (
	"context"
	"fmt"
	"golang-mysql/database"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// func TestExectSql(t *testing.T) {
// 	db := database.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "insert into customer(id,name) values ('arif', 'Arif')"
// 	_, err := db.ExecContext(ctx, script)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Succes Insert Data Into Customer")
// }

// func TestRowsSql(t *testing.T) {
// 	db := database.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "SELECT id, name FROM customer "
// 	rows, err := db.QueryContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id, name string
// 		err := rows.Scan(&id, &name)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("ID :", id)
// 		fmt.Println("NAME :", name)
// 	}
// 	defer rows.Close()
// }

// func TestQuerySqlComplex(t *testing.T) {
// 	db := database.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "select id, name, email, balance, rating, birth_date, created_at from customer"
// 	rows, err := db.QueryContext(ctx, script)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id, name, email string
// 		var balance int32
// 		var rating float64
// 		var birth_date, created_at time.Time

// 		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("==============")
// 		fmt.Println("ID :", id)
// 		fmt.Println("NAME :", name)
// 		fmt.Println("EMAIL :", email)
// 		fmt.Println("BALANCE :", balance)
// 		fmt.Println("RATING :", rating)
// 		fmt.Println("BIRTH_DATE :", birth_date)
// 		fmt.Println("CREATED_AT :", created_at)

// 	}
// 	defer rows.Close()

// }

// func TestSqlInjectionNotSafe(t *testing.T) {

// 	db := database.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	username := "admin'; #"
// 	password := "salah"

// 	script := "SELECT username FROM user where username = '" + username + "' AND password = '" + password + "' LIMIT 1"

// 	rows, err := db.QueryContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		var username string
// 		err := rows.Scan(&username)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Sukses Login", username)
// 	} else {
// 		fmt.Println("Gagal Login", username)
// 	}
// 	defer rows.Close()
// }

// func TestSqlInjectionNotSafe(t *testing.T) {

// 	db := database.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	username := "admin"
// 	password := "admin"

// 	script := "SELECT username FROM user where username = ? AND password = ? LIMIT 1"

// 	rows, err := db.QueryContext(ctx, script, username, password)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		var username string
// 		err := rows.Scan(&username)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Sukses Login", username)
// 	} else {
// 		fmt.Println("Gagal Login", username)
// 	}
// 	defer rows.Close()
// }

// func TestExectSqlSafe(t *testing.T) {
// 	db := database.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	username := "arif"
// 	password := "arif"

// 	script := "INSERT INTO user(username,password) values(?, ?)"
// 	_, err := db.ExecContext(ctx, script, username, password)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Sukses Insert New Customer")
// }

// func TestExectSqlAutoIncreament(t *testing.T) {

// 	db := database.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	email := "arif@gmail.com"
// 	comment := "Test Comment"

// 	script := "insert into comments(email,comment) values (?,?)"

// 	result, err := db.ExecContext(ctx, script, email, comment)
// 	if err != nil {
// 		panic(err)
// 	}

// 	insertID, err := result.LastInsertId()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Success Insert Data To Comments With ID", insertID)
// }

// func TestPrepareStatement(t *testing.T) {

// 	db := database.GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "INSERT INTO comments(email,comment) VALUES(?, ?)"

// 	statement, err := db.PrepareContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer statement.Close()

// 	for i := 0; i < 10; i++ {
// 		email := "holla" + strconv.Itoa(i) + "@gmail.com"
// 		comment := "Test Aja" + strconv.Itoa(i)

// 		result, err := statement.ExecContext(ctx, email, comment)
// 		if err != nil {
// 			panic(err)
// 		}

// 		id, err := result.LastInsertId()
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("Comment Ke", id)
// 	}
// }

func TestTransactionSql(t *testing.T) {

	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	// tx

	for i := 0; i < 10; i++ {
		email := "yollo" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment Aja" + strconv.Itoa(i)

		script := "insert into comments(email, comment) values(?, ?)"
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Insert Comment Ke", id)

	}
	tx.Commit()

}
