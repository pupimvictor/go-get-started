package app

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"strconv"
)

type (
	Item struct {
		Name   string
		Amount int
	}
	ShoppingList struct {
		name  string
		items []Item
	}
)

// our app memory
var myShoppingList *ShoppingList

// init ShoppingList obj to be used as the app memory
func NewShoppingList(name string) {
	myShoppingList = &ShoppingList{
		name:  name,
		items: []Item{},
	}
	log.Printf("Shopping list %s initialized\n", name)
}

// getItems return shoppingList items if any, error if empty
func (sl ShoppingList) getItems() ([]Item, error) {
	if len(sl.items) < 1 {
		return nil, fmt.Errorf("List of \"%s\" is still empty", sl.name)
	}
	return sl.items, nil
}

// addItem to the shoppingList obj
func (sl *ShoppingList) addItem(name string, amount int) {
	i := Item{
		Name:   name,
		Amount: amount,
	}
	sl.items = append(sl.items, i)
	log.Printf("%s added to shopping list\n", name)
}

// GetItemsHandler getItems in shopping list and write response with formatted json message
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	//check shoppingList was instantiated by NewShoppingList
	if myShoppingList == nil {
		w.WriteHeader(500)
		w.Write([]byte("list not initialized"))
		return
	}
	items, err := myShoppingList.getItems()
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	//transform items into a list
	var itemsList []string
	for _, i := range items {
		itemsList = append(itemsList, fmt.Sprintf("%d %s", i.Amount, i.Name))
	}

	//prepare json payload
	jsonItems, err := json.Marshal(itemsList)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	//write http response with json payload
	w.Write(jsonItems)
}

// AddItemHandler decodes request using query parameters and add item to shoppingList
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	//getting value from request url
	name := r.URL.Query().Get("name")
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	//check error when converting to int
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		log.Printf("error %s", err.Error())
		return
	}
	if myShoppingList == nil {
		w.WriteHeader(500)
		w.Write([]byte("list not initialized"))
		return
	}
	myShoppingList.addItem(name, amount)

	//write success status code
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
