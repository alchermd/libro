package main

import (
	"net/http"
	"path/filepath"
	"text/template"
)

// Holds the data that is passed and accessed by the HTML templates.
type templateData struct{}

// Setup a new template cache.
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	// Loads the template files, then the partials, and then the individual page.
	// This keeps the block override working.
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

// Augment the template with a set of default data.
func (app *application) addDefaultData(td *templateData, w http.ResponseWriter, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	return td
}

// Extra functions that the HTML templates can access.
var functions = template.FuncMap{}
