package main

import "net/http"
import "github.com/pupimvictor/go-get-started/sample-app/app"



func main() {
	http.HandleFunc("/item/get", app.GetItemsHandler)
	http.HandleFunc("/item/add", app.AddItemHandler)
	http.ListenAndServe(":8080", nil)
}

