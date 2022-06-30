package router

import (
	"net/http"

	// 3rd party
	"github.com/go-chi/chi/v5"

	// internal
	"github.com/TonyPath/apiserver/router/middleware"
)

type APIRouter struct {
	mux *chi.Mux
	mw  []middleware.Middleware
}

func NewApiRouter() *APIRouter {
	return &APIRouter{
		mux: chi.NewRouter(),
		mw: []middleware.Middleware{
			middleware.Recover(),
		},
	}
}

func (router *APIRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func (router *APIRouter) Handle(method string, path string, handlerFn http.HandlerFunc, mw ...middleware.Middleware) {
	handler := middleware.AddMiddleware(handlerFn, mw)

	handler = middleware.AddMiddleware(handler, router.mw)

	router.mux.Method(method, path, handler)
}
