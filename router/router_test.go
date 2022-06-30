package router

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApiRouter(t *testing.T) {

	router := NewApiRouter()
	router.Handle(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`hello api router`))
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)

	body, _ := ioutil.ReadAll(w.Body)
	require.Equal(t, "hello api router", string(body))
}

func TestApiRouter_Panic(t *testing.T) {

	router := NewApiRouter()
	router.Handle(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		panic("sth terrible happened")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)

	body, _ := ioutil.ReadAll(w.Body)
	require.Equal(t, http.StatusText(http.StatusInternalServerError)+"\n", string(body))
}
