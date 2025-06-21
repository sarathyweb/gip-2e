package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://sarathyweb.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Could not retrive ", url, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not read body")
	}
	fmt.Println(string(body))
}
