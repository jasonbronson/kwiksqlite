package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	//read .env

	// Create listener
	server := http.Server{
		Addr:    ":" + Cfg.Port,
		Handler: NewHTTPHandler(Cfg),
	}

	shutdown := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down gracefully
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout
			log.Printf("HTTP api Shutdown: %v", err)
		}
		close(shutdown)
	}()

	log.Printf("HTTP api listening on: %s", Cfg.Port)
	log.Println(Cfg.DisableSSL)
	if !Cfg.DisableSSL {
		if err := server.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
			// Error starting or closing listener
			log.Printf("HTTPS api ListenAndServeTLS: %v", err)
		}
	} else {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener
			log.Printf("HTTP api ListenAndServe: %v", err)
		}
	}

	// Block until the api gracefully shuts down and finally exit the process
	<-shutdown

}
