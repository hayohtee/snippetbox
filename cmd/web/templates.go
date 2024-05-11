package main

import "github.com/hayohtee/snippetbox/internal/models"

// templateData is a type to act as the holding structure
// for any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Snippet *models.Snippet
}