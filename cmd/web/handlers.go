package main

import "net/http"

// Handles the home page.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.html", nil)
}
