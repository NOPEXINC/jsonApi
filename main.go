// returns a json from an external API
package main

import (
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
		port = "3000"
	}
	http.HandleFunc("/api" / PostsHandler)
	log.Println("listening on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func PostsHandler(res http.ResponseWriter, req *http.Request) {
	url := "https://jsonplaceholder.typicode.com/posts"
	posts := make([]Post, 0)

	GetJSON(url, &posts)

	data := string(json.Marshal(posts))
	log.Println("%s request from %s", req.Method, req.URL.Path)
	fmt.Fprintf(res, "%+v", data)
}
