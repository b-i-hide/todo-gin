package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"time"
)

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Due time.Time `json:"due"`
	Note string `json:"note"`
}


func NewTodo() Todo {
	return Todo{}
}

func (t Todo) GetAll(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query(`SELECT * FROM todos`)
	if err != nil {
		log.Fatalf("error: %s", err)
		return nil, err
	}
	return ScanTodos(rows)
}

func (t Todo) Insert(db *sql.DB, todo Todo) error {
	stmt, err := db.Prepare(`INSERT todos SET title=?, created_at=?, due=?`)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	_, err = stmt.Exec(todo.Title, t.Created_at, todo.Due)
	if err != nil {
		log.Fatalf("error: %s", err)
		return err
	}
	return nil
}

func (t Todo) GetById(db *sql.DB, id int64) (Todo, error) {
	var todo Todo
	err := db.QueryRow("SELECT * FROM todos WHERE id=? LIMIT 1", id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Created_at,
		&todo.Updated_at,
		&todo.Due,
		&todo.Note,
	)
	if err != nil {
		log.Fatalf("error: %s", err)
		return Todo{}, err
	}
	return todo, nil
}

func (t Todo) RemoveById(db *sql.DB, id int64) error {
	stmt, err := db.Prepare(`DELETE FROM todos WHERE id=?`)
	if err != nil {
		log.Fatalf("cannot delete record. error: %s", err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatalf("error: %s", err)
		return err
	}
	return nil
}

func (t Todo) Update(db *sql.DB, todo Todo) error {
	stmt, err := db.Prepare("UPDATE todos SET title=?, due=?, note=? WHERE id=?")
	if err != nil {
		log.Fatalf("cannot update record. error: %s", err)
	}
	_, err = stmt.Exec(todo.Title, todo.Due, todo.Note, todo.Id)
	if err != nil {
		return err
	}
	return nil
}
