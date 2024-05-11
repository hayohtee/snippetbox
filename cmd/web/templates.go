package main

import (
	"html/template"
	"path/filepath"

	"github.com/hayohtee/snippetbox/internal/models"
)

// templateData is a type to act as the holding structure
// for any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

// Create an in memory cache for all the templates needed by the app
func newTemplateCache()(map[string]*template.Template, error) {	
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string {
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}