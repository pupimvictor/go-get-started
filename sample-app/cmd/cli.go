package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	itemName := os.Args[1]
	itemAmount := os.Args[2]

	fmt.Printf("adding %s %s(s) to shopping list\n", itemAmount, itemName)

	url := fmt.Sprintf("http://127.0.0.1:8082/item/add?name=%s&amount=%s", itemName, itemAmount)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s", err)
		log.Fatal(err)
	}
	if resp != nil {
		payload, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%d - %s\n", resp.StatusCode, payload)
	}

	os.Exit(0)
}
