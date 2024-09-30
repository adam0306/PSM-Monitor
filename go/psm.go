package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://<psm>/psm/api/health")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
