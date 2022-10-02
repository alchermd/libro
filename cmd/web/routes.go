package main

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Mapping of the URL routes to their handlers.
func (app *application) routes() http.Handler {
	// Setup middleware chain
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Create a new servemux and apply handler mappings.
	mux := mux.NewRouter()
	mux.HandleFunc("/", app.home).Methods("GET")
	mux.HandleFunc("/books/create", app.createBookForm).Methods("GET")

	// Serve static assets.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	return csrf.Protect([]byte("32-byte-long-auth-key"))(standardMiddleware.Then(mux))
}
