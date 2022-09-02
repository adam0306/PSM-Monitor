package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://example.com")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
