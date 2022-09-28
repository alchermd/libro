package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	// Load CLI options
	serverPort := flag.String("addr", ":4000", "Port that the server runs on")
	flag.Parse()

	// Setup custom loggers.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Setup template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: templateCache,
	}

	svr := &http.Server{
		Addr:         *serverPort,
		ErrorLog:     app.errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.infoLog.Printf("Starting server on %s", svr.Addr)
	app.errorLog.Fatal(svr.ListenAndServe())
}
