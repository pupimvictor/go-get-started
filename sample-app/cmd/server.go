package main

import (
	"log"
	"net/http"
	"os"
)
import "github.com/pupimvictor/go-get-started/sample-app"

func main() {
	shoppingListName := os.Args[1]

	app.NewShoppingList(shoppingListName)
	log.Println("listening on port 5001")

	http.HandleFunc("/item/get", app.GetItemsHandler)
	http.HandleFunc("/item/add", app.AddItemHandler)
	log.Fatal(http.ListenAndServe(":5001", nil))
}
