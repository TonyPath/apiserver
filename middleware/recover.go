package middleware

import "net/http"

func Recover() MidFunc {

	mw := func(handler http.Handler) http.Handler {

		h := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			handler.ServeHTTP(w, r)
		}

		return http.HandlerFunc(h)
	}

	return mw
}
