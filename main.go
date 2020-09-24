package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	
	// Concurrent Serving while waiting for a kill signal
	go func(){
		err := s.ListenAndServe()
		if(err!=nil){
			l.Fatal(err)
		}
	}();
	
	// Map the routes in servemux and start server
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// Graceful Exit - shutdown
	sigChan := make(chan os.Signal)

	// Notify the channel when an interrupt or a kill signal is encountered
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	// Extract the value of the channel (Blocking Operation)
	sig := <-sigChan
	l.Println("Received a termination request :: ", sig)

	// Get Context(Timeout Method) and shutdown
	c, _ := context.WithTimeout(context.TODO(), 30*time.Second)
	s.Shutdown(c)
}
