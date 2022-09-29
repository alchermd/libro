package psql

import (
	"database/sql"
	"time"
)

type BookModel struct {
	DB *sql.DB
}

// Insert a new book into the database, returning its ID.
func (m *BookModel) Insert(title, description, coverPhotoURL, author string, publishedAt time.Time, pages int, genres []string) (int, error) {
	return 0, nil
}
