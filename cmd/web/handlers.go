package main

import "net/http"

// Handles the home page.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Temporary code to insert genres
	// TODO: Delete me!
	id, err := app.genres.Insert("test-genre")
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("Inserted genre ID: %d", id)

	app.render(w, r, "home.page.html", nil)
}
