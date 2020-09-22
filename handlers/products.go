package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)
// Creating a type to receive Global Logger
type Hello struct {
	l *log.Logger
}

// Convert Logger type into Hello Type
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Implementing the ServeHTTP Method for Handler Interface
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Read data from request body
	d, err := ioutil.ReadAll(r.Body)
	// Handle the error
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("Error Encountered"))
		// OR
		http.Error(w, "Error Encountered!!\n", http.StatusBadRequest)
		return
	}
	// Log the data
	h.l.Println(string(d))
}
