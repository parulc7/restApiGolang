package main

import (
	"time"
	"log"
	"net/http"
	"os"

	"github.com/parulc7/CoffeeShopAPI/handlers"
)

func main() {
	// Global Uniform Logger
	l := log.New(os.Stdout, "coffee-shop", log.LstdFlags)
	// Passing the global logger to the handler
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	// Create a custom servemux
	sm := http.NewServeMux()

	// Server Config here
	s:= &http.Server{
		Addr:"",
		Handler:sm,
		IdleTimeout:120*time.Second,
		ReadTimeout:1*time.Second,
		WriteTimeout:1*time.Second
	}
	// Map the routes in servemux and start server
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	s.ListenAndServe(":8080", sm)
}
