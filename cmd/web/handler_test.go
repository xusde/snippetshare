package main

import (
	"net/http"

	"testing"

	"github.com/xusde/snippetshare/internal/assert"
)

func TestPing(t *testing.T) {
	// Create a new instance of our application
	app := newTestApplication(t)

	// Establish a new test server for running end-to-end tests.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Make a GET /ping request and check that we get back a 200 OK response.
	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}
