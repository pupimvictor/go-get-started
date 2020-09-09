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
	itemName := os.Args[1]
	itemAmount := os.Args[2]

	fmt.Printf("adding %s %s(s) to shopping list\n", itemAmount, itemName)

	//prepare URL
	url := fmt.Sprintf("http://127.0.0.1:5001/item/add?name=%s&amount=%s", itemName, itemAmount)

	// HTTP Get Request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s", err)
		//log error to stdout and exit program with error code
		log.Fatal(err)
	}
	if resp != nil {
		//Read response payload
		payload, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%d - %s\n", resp.StatusCode, payload)
	}

	//exit program
	os.Exit(0)
}
