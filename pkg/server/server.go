package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/ssengalanto/biscuit/pkg/constants"
)

type Server struct {
	server *http.Server
}

// New creates a new HTTP server.
func New(addr int, mux *chi.Mux) *Server {
	srv := &Server{
		server: &http.Server{
			Addr:         ":" + strconv.Itoa(addr),
			Handler:      mux,
			IdleTimeout:  constants.IdleTimeout,
			ReadTimeout:  constants.ReadTimeout,
			WriteTimeout: constants.WriteTimeout,
		},
	}

	return srv
}

// Start starts the HTTP server with graceful shutdown.
func (s *Server) Start() error {
	errStream := gracefulShutdown(s.server)

	err := s.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return ErrServerClosed
	}

	err = <-errStream
	if err != nil {
		return err
	}

	return nil
}

// gracefulShutdown gracefully shuts down the server.
func gracefulShutdown(s *http.Server) <-chan error {
	errStream := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), constants.ShutdownGracePeriod)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			errStream <- err
		}

		errStream <- nil
	}()

	return errStream
}
