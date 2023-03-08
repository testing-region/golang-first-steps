package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	robots, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
