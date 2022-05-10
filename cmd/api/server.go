package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Channel for receiving errors returned by the graceful Shutdown()
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		// Listen for incoming SIGINT and SIGTERM. Relay them to quit channel
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		app.logger.PrintInfo("shutting down server", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		shutdownError <- srv.Shutdown(ctx)
	}()

	app.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})

	// Calling Shutdown() on our server will cause ListenAndServe() to immediately
	// return a http.ErrServerClosed error. If that is the case, it means that the
	// graceful shutdown has started. So we check if it is NOT http.ErrServerClosed.
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	// If returned value is an error, there was a problem with the graceful shutdown.
	err = <-shutdownError
	if err != nil {
		return err
	}

	// At this point the graceful shutdown completed successfully.
	app.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
