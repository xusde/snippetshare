package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/xusde/snippetshare/internal/models/mocks"
)

// Create a newTestApplication helper which returns an instance of our
// application struct containing mocked dependencies.
func newTestApplication(t *testing.T) *application {
	// Create an instance of the template cache.
	templateCache, err := newTemplateCache()
	if err != nil {
		t.Fatal(err)
	}

	// And a form decoder.
	formDecoder := form.NewDecoder()

	// And a session manager instance. Note that we use the same settings as
	// production, except that we *don't* set a Store for the session manager.
	// If no store is set, the SCS package will default to using a transient
	// in-memory store, which is ideal for testing purposes.
	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	return &application{
		logger:         slog.New(slog.NewTextHandler(io.Discard, nil)),
		snippets:       &mocks.SnippetModel{}, // Use the mock.
		users:          &mocks.UserModel{},    // Use the mock.
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}
}

// Define a custom testServer type which anonymously embeds a httptest.Server
// instance.
type testServer struct {
	*httptest.Server
}

// Create a newTestServer helper which initializes and returns a new instance
// of our custom testServer struct.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	// Initialize a new httptest.Server, passing in the httprouter.Handler
	// as the "handler" parameter.
	ts := httptest.NewTLSServer(h)

	// Initialize a new cookie jar.
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	// Add the cookie jar to the client, so that response cookies are stored
	// and then sent with subsequent requests.
	ts.Client().Jar = jar
	// Disable redirect-following for the client. By default, the client will
	// automatically follow any redirects and update the request accordingly.
	ts.Client().CheckRedirect = func(r *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// Initialize a new testServer embedding the httptest.Server instance.
	return &testServer{ts}
}

// Implement a get method on our custom testServer type. This makes a GET
// request to a given URL path on the test server, and returns the response
// status code, headers, and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	return rs.StatusCode, rs.Header, string(body)
}
