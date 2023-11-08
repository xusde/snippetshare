package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	Id      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?
	`
	row := m.DB.QueryRow(stmt, id)
	var s Snippet
	err := row.Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}
	return s, nil
}

// return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	// write the sql statement we want to execute
	stmt := `select id, title, content, created, expires from snippets
	where expires > UTC_TIMESTAMP() order by id desc limit 10`

	// use the Query() method on the connection pool to execute our sql statement
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// make sure to close the rows when we leave
	defer rows.Close()

	// initialize empty slice to hold the snippets
	var snippets []Snippet
	for rows.Next() {

		// create a pointer to a new zeroed Snippet struct
		var s Snippet
		err = rows.Scan(&s.Id, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	// return snippets slice if everything is okay
	return snippets, nil
}
