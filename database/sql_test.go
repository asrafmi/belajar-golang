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
