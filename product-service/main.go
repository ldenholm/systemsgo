package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ldenholm/systemsgo/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	product := handlers.NewProducts(l)

	// ServeMux
	sm := http.NewServeMux()

	// Assign Handlers
	sm.Handle("/", helloHandler)
	sm.Handle("/products", product)

	// Create Server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Signal Channel
	// notify on os commands (interrupt & kill)
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	l.Println("Shutdown received, gracefully shutting down", sig)

	timeout, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// Graceful shutdown
	server.Shutdown(timeout)
}
