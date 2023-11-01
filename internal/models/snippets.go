package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires time.Time) (int, error) {
	return 0, nil
}
// return a specific snippet based on its id
func (m *SnippetModel) Get (id int ) (Snippet,error) {
	return Snippet{}, nil
}
// return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}