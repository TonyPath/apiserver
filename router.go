package apiserver

import (
	"github.com/TonyPath/apiserver/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	mux *chi.Mux
	mw  []middleware.Middleware
}

func NewRouter() *Router {
	return &Router{
		mux: chi.NewRouter(),
		mw: []middleware.Middleware{
			middleware.Recover(),
		},
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func (router *Router) Handle(method string, path string, handlerFn http.HandlerFunc, mw ...middleware.Middleware) {
	handler := middleware.AddMiddleware(handlerFn, mw)

	handler = middleware.AddMiddleware(handler, router.mw)

	router.mux.Method(method, path, handler)
}
