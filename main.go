package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/shubash/saibaba/router"
)

func main() {
	r := router.Router()

	// Create a new CORS handler with the desired options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allowing all origins
		AllowCredentials: true,
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(r)

	port := "8081" // Change the port to your desired value

	fmt.Printf("Server is ready and running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
