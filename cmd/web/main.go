package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/alchermd/aklatan/internals/models/psql"
	_ "github.com/lib/pq"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	genres        psql.GenreModel
}

func main() {
	// Load CLI options
	serverPort := flag.String("addr", ":4000", "Port that the server runs on")
	dbHost := flag.String("dbHost", "localhost", "Database host")
	dbPort := flag.String("dbPort", "5432", "Database port")
	dbUser := flag.String("dbUser", "libro_user", "Database username")
	dbPassword := flag.String("dbPassword", ":libro_pass", "Database password")
	dbName := flag.String("dbName", "libro", "Database name")
	flag.Parse()

	// Setup custom loggers.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Setup template cache
	infoLog.Print("Setting up template cache.")
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Print("Setting up template cache -- success!")

	// Setup database.
	infoLog.Print("Setting up database.")
	db, err := openDB(*dbHost, *dbPort, *dbUser, *dbPassword, *dbName)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	infoLog.Print("Setting up database -- success!")

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: templateCache,
		genres:        psql.GenreModel{DB: db},
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

// Helper function for initial database setup.
func openDB(host, port, user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
