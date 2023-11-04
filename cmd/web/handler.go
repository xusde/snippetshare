package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		app.notfound(w)
		return
	}
	files := []string {
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		app.serverError(w, r, err)
	}
}

func (app *application)snippetView(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notfound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application)snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}


	// create data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// pass data to SnippetModel.Insert
	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// Redirect
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

}