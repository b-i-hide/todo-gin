package models

import (
	"database/sql"
)

func ScanTodos(r *sql.Rows) ([]Todo, error) {
	var s Todo
	var err error
	todos := make([]Todo, 0)
	for r.Next() {
		err = r.Scan(
			&s.Id,
			&s.Title,
			&s.Created_at,
			&s.Updated_at,
			&s.Due,
			&s.Note,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, s)
	}
	return todos, nil
}
