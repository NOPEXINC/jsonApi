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
	http.HandleFunc("/api", PostsHandler)
	log.Println("listening on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func PostsHandler(res http.ResponseWriter, req *http.Request) {
	url := "https://jsonplaceholder.typicode.com/posts"
	posts := make([]Post, 0)

	GetJSON(url, &posts)

	data, _ := json.Marshal(posts)

	log.Printf("getting %s requests from %s", req.Method, req.URL)
	fmt.Fprintf(res, "%+v", string(data))
}

func GetJSON(url string, target interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
