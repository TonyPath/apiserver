package middleware

import "net/http"

type MidFunc func(next http.Handler) http.Handler

func Wrap(handler http.Handler, mw ...MidFunc) http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}
