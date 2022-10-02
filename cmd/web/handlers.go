package main

import (
	"net/http"
	"time"
)

// Handles the home page.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Temporary code to insert a book
	// TODO: Delete me!
	publishedAt, _ := time.Parse("02/01/2006", "01/01/2013")
	app.infoLog.Print(publishedAt)
	id, err := app.books.InsertBook(
		"The Phoenix Project",
		"Bill is an IT manager at Parts Unlimited. It's Tuesday morning and on his drive into the office, Bill gets a call from the CEO.",
		"https://kbimages1-a.akamaihd.net/65bdc1f3-a7c3-4d14-8736-45a26dc7911d/1200/1200/False/the-phoenix-project-5th-anniversary-edition.jpg",
		"Gene Kim",
		publishedAt,
		345,
		[]string{
			"Business",
			"Technology",
			"Fiction",
			"Management",
		},
	)

	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("Inserted book ID: %d", id)

	app.render(w, r, "home.page.html", nil)
}

// Displays the book creation form.
func (app *application) createBookForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create-book.page.html", nil)
}
