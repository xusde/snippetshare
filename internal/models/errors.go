package models

import (
	"errors"
)

var (
	// If no matching record snippet is found in database
	ErrNoRecord = errors.New("models: no matching record found")
	// If user tries to login in with inccorrect email or password
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// If user tries to sign up with duplicate email
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
