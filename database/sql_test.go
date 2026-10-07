package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	q := "INSERT INTO customer(id, name) VALUES('prbw', 'prabowo')"
	_, err := db.ExecContext(ctx, q)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	q := "SELECT id, name FROM customer;"
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id => ", id)
		fmt.Println("Name => ", name)
	}

	fmt.Println("Success")
}

func TestComplexQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	q := "SELECT id, name, email, balance, rating, birth_date, married FROM customer;"
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int64
		var rating float64
		var birth_date sql.NullTime
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("============")
		fmt.Println("Id => ", id)
		fmt.Println("Name => ", name)
		if email.Valid {
			fmt.Println("Email => ", email.String)
		}
		fmt.Println("Balance => ", balance)
		fmt.Println("Rating => ", rating)
		if birth_date.Valid {
			fmt.Println("Birth Date => ", birth_date.Time)
		}
		fmt.Println("Married => ", married)
	}

	fmt.Println("Success")
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "jkw'; #"
	password := "passwordsalah"

	ctx := context.Background()
	q := "SELECT username, password FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1;"
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username, password string
		err = rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses login")
		fmt.Println("Username => ", username)
		fmt.Println("Password => ", password)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "jkw'; #"
	password := "passwordsalah"

	ctx := context.Background()
	q := "SELECT username, password FROM users WHERE username = ? AND password = ? LIMIT 1;"
	rows, err := db.QueryContext(ctx, q, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username, password string
		err = rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses login")
		fmt.Println("Username => ", username)
		fmt.Println("Password => ", password)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestExecSQLSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	q := "INSERT INTO customer(id, name) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, q, "prbw", "prabowo")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}

func TestExecLastInsertID(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "prabowo@example.com"
	comment := "prabowo gemoy"

	ctx := context.Background()
	q := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, q, email, comment)
	if err != nil {
		panic(err)
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", lastInsertedId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	q := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, q)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("user%d@example.com", i)
		comment := fmt.Sprintf("Comment %d", i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertedId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert new comment with id", lastInsertedId)
	}
}
