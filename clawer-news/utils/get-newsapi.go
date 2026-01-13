package utils

import (
	"io"
	"log"
	"net/http"
)

func ConnectAPI() ([]byte, error) {
	apiKey := "589bba4de4bf40949e6fedd51692f61f"
	url := "https://newsapi.org/v2/top-headlines?country=us&apiKey=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Request failed: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return body, err

}
