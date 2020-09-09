package app

import (
	"encoding/json"
	"fmt"

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

var myShoppingList *ShoppingList

func NewShoppingList(name string) {
	myShoppingList = &ShoppingList{
		name:  name,
		items: []Item{},
	}
}

func (sl ShoppingList) getItems() ([]Item, error) {
	if len(sl.items) < 1 {
		return nil, fmt.Errorf("empty list", nil)
	}
	return sl.items, nil
}

func (sl *ShoppingList) addItem(name string, amount int) {
	i := Item{
		Name:   name,
		Amount: amount,
	}
	sl.items = append(sl.items, i)
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	if myShoppingList == nil {
		w.Write([]byte("list not initialized"))
	}
	items, err := myShoppingList.getItems()
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}
	jsonItems, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(jsonItems)
}
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	myShoppingList.addItem(name, amount)
}
