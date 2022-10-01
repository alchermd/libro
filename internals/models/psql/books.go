package psql

import (
	"database/sql"
	_ "embed"
	"time"
)

type BookModel struct {
	DB *sql.DB
}

var (
	//go:embed sql/insert-book.sql
	insertBookStmt string
	//go:embed sql/insert-genre.sql
	insertGenreStmt string
	//go:embed sql/insert-book-genre.sql
	insertBookGenreStmt string
)

// Insert a new book into the database, returning its ID.
func (m *BookModel) InsertBook(title, description, coverPhotoURL, author string, publishedAt time.Time, pages int, genres []string) (int, error) {
	bookId := 0
	err := m.DB.QueryRow(insertBookStmt, title, description, coverPhotoURL, author, publishedAt, pages).Scan(&bookId)

	for _, genre := range genres {
		genreId, err := m.InsertGenre(genre)
		if err != nil {
			return 0, err
		}

		if err = m.InsertBookGenre(bookId, genreId); err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	return bookId, nil
}

// Insert a new genre into the database, returning its ID.
func (m *BookModel) InsertGenre(name string) (int, error) {
	lastInsertId := 0
	err := m.DB.QueryRow(insertGenreStmt, name).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

// Insert a mapping for a book and its genre into the database.
func (m *BookModel) InsertBookGenre(bookId, genreId int) error {
	_, err := m.DB.Exec(insertBookGenreStmt, bookId, genreId)
	return err
}
