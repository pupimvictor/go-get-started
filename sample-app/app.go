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
	fmt.Printf("Shopping list %s initialized\n", name)
}

func (sl ShoppingList) getItems() ([]Item, error) {
	if len(sl.items) < 1 {
		return nil, fmt.Errorf( "%s is empty", sl.name)
	}
	return sl.items, nil
}

func (sl *ShoppingList) addItem(name string, amount int) {
	i := Item{
		Name:   name,
		Amount: amount,
	}
	sl.items = append(sl.items, i)
	fmt.Printf("%s added to shopping list\n", name)
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
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

	var itemsList []string
	for _, i := range items {
		itemsList = append(itemsList, fmt.Sprintf("%d %s", i.Amount, i.Name))
	}

	jsonItems, err := json.Marshal(itemsList)
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
		fmt.Printf("error %s", err.Error())
		return
	}
	if myShoppingList == nil {
		w.WriteHeader(500)
		w.Write([]byte("list not initialized"))
		return
	}
	myShoppingList.addItem(name, amount)
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
