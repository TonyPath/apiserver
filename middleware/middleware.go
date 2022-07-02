package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

func AddMiddleware(handler http.Handler, mw []Middleware) http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}
