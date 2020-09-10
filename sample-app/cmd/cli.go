package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	//get os args from execution params
	if len(os.Args) < 2 {
		log.Println("Usage: go-cli <item> <amount>")
		os.Exit(1)
	}
	itemName := os.Args[1]
	itemAmount := os.Args[2]

	log.Printf("adding %s %s(s) to shopping list\n", itemAmount, itemName)

	//prepare URL
	url := fmt.Sprintf("http://127.0.0.1:5001/item/add?name=%s&amount=%s", itemName, itemAmount)

	// HTTP Get Request
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error: %s", err)
		//log error to stdout and exit program with error code
		log.Fatal(err)
	}
	if resp != nil {
		//Read response payload
		payload, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%d - %s\n", resp.StatusCode, payload)
	}

	//exit program
	os.Exit(0)
}
