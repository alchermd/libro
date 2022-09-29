package psql

import (
	"database/sql"
	_ "embed"
)

type GenreModel struct {
	DB *sql.DB
}

var (
	//go:embed sql/insert-genre.sql
	insertGenreStmt string
)

// Insert a new genre into the database, returning its ID.
func (m *GenreModel) Insert(name string) (int, error) {
	lastInsertId := 0
	err := m.DB.QueryRow(insertGenreStmt, name).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}
