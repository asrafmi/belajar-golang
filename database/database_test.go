package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/app_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
}
