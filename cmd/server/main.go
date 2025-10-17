package main

import (
	"fmt"
	"net/http"

	"unit-converter/internal/handler" // Path assumes module name is unit-converter
)

// main sets up the router and starts the server.
func main() {
	// Set up the API endpoint
	// Use the exported ConversionHandler from the handler package
	http.HandleFunc("/api/v1/convert", handler.ConversionHandler)

	port := ":8080"
	fmt.Printf("Starting Unit Converter API server on port %s\n", port)
	
	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
