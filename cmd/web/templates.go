package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/xusde/snippetshare/internal/models"
	"github.com/xusde/snippetshare/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	return t.Format("Jan 02 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// Use the fs.Glob function to get a slice of all filepaths with
	// the pattern './ui/html/pages/*.tmpl'. This gives us a slice of all
	// filepaths for our application 'page' templates
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		// // Add paritial templates into the template set
		// ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		// if err != nil {
		// 	return nil, err
		// }
		// // Add the page template into the template set
		// ts, err = ts.ParseFiles(page)
		// if err != nil {
		// 	return nil, err
		// }
		// Add the template set to the cache, using the name of the page as the key
		cache[name] = ts
	}

	//  Return the map
	return cache, nil
}
