// returns a json from an external API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/api" / PostsHandler)
	log.Println("listening on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
