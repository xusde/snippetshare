package middeware

import "net/http"

func myMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
