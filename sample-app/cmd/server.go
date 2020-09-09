package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
import "github.com/pupimvictor/go-get-started/sample-app"



func main() {
	shoppingListName := os.Args[1]

	app.NewShoppingList(shoppingListName)
	fmt.Println("listening on port 8082")

	http.HandleFunc("/item/get", app.GetItemsHandler)
	http.HandleFunc("/item/add", app.AddItemHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

