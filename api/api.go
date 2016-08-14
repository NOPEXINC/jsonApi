package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// fetches json data from url and then parses
// it into the target object. this function returns  an error
func GetJSON(url string, target interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// parses json data into the the target object
// returns either a byte slice or an error
func GetJsonResponse(target interface{}) ([]byte, error) {
	return json.MarshalIndent(target, "", "  ")
}
