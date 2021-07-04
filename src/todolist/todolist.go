package todolist

import (
	"errors"
	"fmt"
)

type CreateItem struct {
	text string
}

// structs

type Item struct {
	text      string
	id        string
	isChecked bool
}

type Todolist struct {
	store []Item
}

func (todolist *Todolist) addItem(createItem CreateItem) (Item, error) {
	item := Item{
		id:        "",
		text:      createItem.text,
		isChecked: false,
	}

	todolist.store = append(todolist.store, item)

	return item, errors.New("")
}

func (todolist *Todolist) removeItem(item Item) (Item, error) {
	var result Item

	for i := 0; i < len(todolist.store); i++ {
		if todolist.store[i].id == item.id {
			result = todolist.store[i]
			return result, nil
		}
	}

	return result, fmt.Errorf("item with id %s was not found", item.id)
}

func TodoList(item Item) bool {

	return true
}
