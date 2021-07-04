package todolist

import (
	"errors"
)

//
type Item struct {
	id        string
	text      string
	isChecked bool
}

type Todolist struct {
	store []Item
}

func (todolist *Todolist) addItem(item Item) (Item, error, ) {
	storeBefore := todolist.store

	todolist.store = append(todolist.store, item)

	storeAfter := todolist.store

	if len(storeBefore) < len(storeAfter) {
		return item, nil
	} else {
		return Item{
			id:        generateID(),
			text:      "nil",
			isChecked: false,
		}, errors.New("incorrect data")
	}
}

func (todolist *Todolist)removeItem(item Item) (Item, error){

	var result Item

	for i := 0; i < len(todolist.store); i++ {
		if todolist.store[i].id == item.id{
			result = todolist.store[i]
			return result, nil
		}
	}

	return result, errors.New("")
}
func TodoList(item Item) bool {

	return true
}
