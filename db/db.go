package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitDB() (*sql.DB) {
	db, err := sql.Open("mysql", "root@/todo_gin_development?parseTime=true")
	if err != nil {
		log.Fatalf("cannot open database. error: %s", err)
		return nil
	}
	return db
}
