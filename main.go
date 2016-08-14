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

type Tour struct {
	Title  string `json:"packageTitle"`
	Name   string `json:"name"`
	Blurb  string `json:"blurb"`
	Price  string `json:"price"`
	Region string `json:"region"`
}

type Food struct {
	Price Price `json:"prices"`
}

type Price struct {
	Fruits     Fruit     `json:"fruits"`
	Vegetables Vegetable `json:"vegetables"`
}

type Fruit map[string]int
type Vegetable map[string]int

var PORT string = os.Getenv("PORT")

func main() {
	if PORT == "" {
		PORT = "4000"
	}
	http.HandleFunc("/api", PostsHandler)
	http.HandleFunc("/prices", PriceHandler)
	http.HandleFunc("/tours", ToursHandler)
	log.Println("listening on http://localhost:" + PORT)
	http.ListenAndServe(":"+PORT, nil)
}

func PostsHandler(res http.ResponseWriter, req *http.Request) {
	url := "https://jsonplaceholder.typicode.com/posts"
	posts := make([]Post, 0)

	if err := api.GetJSON(url, &posts); err != nil {
		log.Fatal(err)
	}

	data, _ := json.MarshalIndent(posts, "", "   ")

	res.Header().Set("Content-Type", "text/json")
	log.Printf("accepting %s requests from http://localhost:%s%s", req.Method, PORT, req.URL)
	fmt.Fprintf(res, "%+v", string(data))
}

func ToursHandler(res http.ResponseWriter, req *http.Request) {
	url := "http://services.explorecalifornia.org/json/tours.php"
	tours := make([]Tour, 0)

	if err := api.GetJSON(url, &tours); err != nil {
		log.Fatal(err)
	}
	data, _ := json.MarshalIndent(tours, "", "  ")
	res.Header().Set("Content-Type", "text/json")
	log.Printf("accepting %s requests from http://localhost:%s%s", req.Method, PORT, req.URL)
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

	res.Header().Set("Content-Type", "text/json")
	log.Printf("accepting %s requests from http://localhost:%s%s", req.Method, PORT, req.URL)
	fmt.Fprintf(res, "%+v", string(response))
}
