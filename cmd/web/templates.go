package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/hayohtee/snippetbox/internal/models"
)

// templateData is a type to act as the holding structure
// for any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
}

// Create an in memory cache for all the templates needed by the app
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
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

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanData(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanData,
}
