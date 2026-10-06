package database

import (
	"context"
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
