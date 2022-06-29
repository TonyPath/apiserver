package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/TonyPath/apiserver/middleware"
)

const (
	httpPort            = 50000
	httpReadTimeout     = 30 * time.Second
	httpWriteTimeout    = 60 * time.Second
	shutdownGracePeriod = 5 * time.Second
)

func newDefaultConfig() *Config {
	return &Config{
		port:                httpPort,
		readTimeout:         httpReadTimeout,
		writeTimeout:        httpWriteTimeout,
		shutdownGracePeriod: shutdownGracePeriod,
	}
}

type ApiServer struct {
	config  *Config
	mux     *chi.Mux
	httpSrv *http.Server
}

func New(opts ...Option) (*ApiServer, error) {
	cfg := newDefaultConfig()

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	apiSrv := ApiServer{
		mux:    chi.NewRouter(),
		config: cfg,
	}

	return &apiSrv, nil
}

func (apiSrv *ApiServer) Handle(method string, path string, handler http.HandlerFunc) {
	hdl := middleware.Wrap(handler, middleware.Recover())

	apiSrv.mux.Method(method, path, hdl)
}

func (apiSrv *ApiServer) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", apiSrv.config.port),
		ReadTimeout:  apiSrv.config.readTimeout,
		WriteTimeout: apiSrv.config.writeTimeout,
		Handler:      apiSrv.mux,
	}
	apiSrv.httpSrv = srv

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		if err := apiSrv.shutdown(ctx); err != nil {
			return err
		}
		return ctx.Err()
	case err := <-serverErrors:
		return err
	}
}

func (apiSrv *ApiServer) shutdown(ctx context.Context) error {
	tctx, cancel := context.WithTimeout(context.Background(), shutdownGracePeriod)
	defer cancel()

	if err := apiSrv.httpSrv.Shutdown(tctx); err != nil {
		_ = apiSrv.httpSrv.Close()
		return fmt.Errorf("cannot stop api server gracefully: %w", err)
	}

	return nil
}
