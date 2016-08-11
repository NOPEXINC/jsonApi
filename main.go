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

type Food struct {
	Price Price
}

type Price struct {
	Fruits     Fruit
	Vegetables Vegetable
}

type Fruit map[string]int
type Vegetable map[string]int

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	http.HandleFunc("/api", PostsHandler)
	http.HandleFunc("/prices", PriceHandler)
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

func PriceHandler(res http.ResponseWriter, req *http.Request) {
	fruits := make(map[string]int)
	veggies := make(map[string]int)

	fruits["Bananas"] = 1000
	fruits["Mangoes"] = 200
	fruits["Pineaples"] = 2400

	veggies["Tomatoes"] = 200
	veggies["Beans"] = 1500
	veggies["Onions"] = 350

	price := Price{fruits, veggies}
	food := Food{price}

	response, err := api.GetJsonResponse(&food)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("accepting %s requests from http://localhost:4000%s", req.Method, req.URL)
	fmt.Fprintf(res, "%+v", string(response))
}
