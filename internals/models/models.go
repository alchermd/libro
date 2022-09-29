package models

import "time"

// Represents a book.
type Book struct {
	Id            int
	Title         string
	Description   string
	CoverPhotoURL string
	Author        string
	PublishedAt   time.Time
	Pages         int
	Genres        []Genre
}

// Represents a genre of a book.
type Genre struct {
	Id   int
	Name string
}
