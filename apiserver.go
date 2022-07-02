package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	httpPort            = 8080
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

type APIServer struct {
	config  *Config
	router  *Router
	httpSrv *http.Server
}

func New(router *Router, opts ...ServerConfig) (*APIServer, error) {
	cfg := newDefaultConfig()

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	apiSrv := APIServer{
		router: router,
		config: cfg,
	}

	return &apiSrv, nil
}

func (apiSrv *APIServer) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", apiSrv.config.port),
		ReadTimeout:  apiSrv.config.readTimeout,
		WriteTimeout: apiSrv.config.writeTimeout,
		Handler:      apiSrv.router,
	}
	apiSrv.httpSrv = srv

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		if err := apiSrv.Shutdown(); err != nil {
			return err
		}
		return ctx.Err()
	case err := <-serverErrors:
		return err
	}
}

func (apiSrv *APIServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownGracePeriod)
	defer cancel()

	if err := apiSrv.httpSrv.Shutdown(ctx); err != nil {
		_ = apiSrv.httpSrv.Close()
		return fmt.Errorf("cannot stop api server gracefully: %w", err)
	}

	return nil
}
