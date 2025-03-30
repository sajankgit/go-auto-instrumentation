package main

import (
	"fmt"
	"net/http"
)

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func main() {
	// Registering the handlers with their respective paths
	http.HandleFunc("/", homeHandler)

	// Starting the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
