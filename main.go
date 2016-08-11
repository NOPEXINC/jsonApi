// returns a json from an external API
package main

import (
	"./api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Post struct {
	UserId int    `json:"userId"`
	PostId int    `json::"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	http.HandleFunc("/api", PostsHandler)
	log.Println("listening on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func PostsHandler(res http.ResponseWriter, req *http.Request) {
	url := "https://jsonplaceholder.typicode.com/posts"
	posts := make([]Post, 0)

	api.GetJSON(url, &posts)

	data, _ := json.MarshalIndent(posts, "", "   ")

	log.Printf("accepting %s requests from http://localhost:4000%s", req.Method, req.URL)
	fmt.Fprintf(res, "%+v", string(data))
}
