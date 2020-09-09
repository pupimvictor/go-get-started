package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"os"
)

func main() {
	itemName := os.Args[1]
	itemAmount := os.Args[2]

	url := fmt.Sprintf("127.0.0.1:8080/items/add?name=%s&amount=%s", itemName, itemAmount)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %v - %d", err, resp.StatusCode)
		log.Fatal(err)
	}

	fmt.Printf("%s added to shopping list", itemName)
	os.Exit(0)
}
