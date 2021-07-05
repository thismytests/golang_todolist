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

func (todolist Todolist) findItemByIdInStore(id string) (Item, error) {
	for i := 0; i < len(todolist.store); i++ {
		if todolist.store[i].id == id {
			return todolist.store[i], nil
		}
	}

	return Item{}, errors.New("Not found")
}

func (todolist *Todolist) addItem(createItem CreateItem) (Item, error) {
	item := Item{
		id:        generateID(),
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

func (todolist *Todolist) updateItem(id string, item Item) (Item, error) {
	errMessage := fmt.Sprintf("updatedItem with id %s was not founded", item.id)

	for i := 0; i < len(todolist.store); i++ {
		if todolist.store[i].id == item.id {
			todolist.store[i] = item
			return item, nil
		}
	}

	return Item{}, errors.New(errMessage)

}
