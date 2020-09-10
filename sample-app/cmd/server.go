package main

import (
	"log"
	"net/http"
	"os"
)
import "github.com/pupimvictor/go-get-started/sample-app"

func main() {
	//get os args from execution params
	if len(os.Args) < 2 {
		log.Println("Usage: server <List Name>")
		os.Exit(1)
	}
	shoppingListName := os.Args[1]

	app.NewShoppingList(shoppingListName)
	log.Println("listening on port 5001")

	http.HandleFunc("/item/get", app.GetItemsHandler)
	http.HandleFunc("/item/add", app.AddItemHandler)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("hello, friend"))
	})

	log.Fatal(http.ListenAndServe(":5001", nil))
}
