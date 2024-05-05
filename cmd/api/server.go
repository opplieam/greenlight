package main

import (
	"context"
	"errors"
	"fmt"
	"log"
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
		ErrorLog:     log.New(app.logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	shutdownError := make(chan error)

	// start background goroutine
	go func() {
		// create quit channel which carries os.Signal values
		// buffered channel with size 1
		quit := make(chan os.Signal, 1)
		// Notify() listen for incoming SIGINT and SIGTERM signal
		// and relay to quit channel.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		// Read the signal from quit channel. This code will block until a signal is received
		s := <-quit

		app.logger.PrintInfo("shutting down server", map[string]string{
			"signal": s.String(),
		})

		// 20 sec to complete remaining task before shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		// Call Shutdown() on the server like before, but now we only send on the
		// shutdownError channel if it returns an error.
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}
		// Log a message to say that we're waiting for any background goroutines to
		// complete their tasks.
		app.logger.PrintInfo("completing background tasks", map[string]string{
			"addr": srv.Addr})
		// Call Wait() to block until our WaitGroup counter is zero --- essentially
		// blocking until the background goroutines have finished. Then we return nil on
		// the shutdownError channel, to indicate that the shutdown completed without
		// any issues.
		app.wg.Wait()
		shutdownError <- nil
	}()

	app.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}
	app.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})
	return nil
}
