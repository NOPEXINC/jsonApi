package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetJSON(url string, target interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetJsonResponse(target interface{}) ([]byte, error) {
	return json.MarshalIndent(target, "", "  ")
}
