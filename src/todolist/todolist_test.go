package todolist

import (
	"fmt"
	"testing"
)

func TestTodoListAddItem(t *testing.T) {
	createItem := CreateItem{
		text: "text",
	}

	todolist := Todolist{store: []Item{}}

	todolistLenStoreBefore := len(todolist.store)

	todolist.addItem(createItem)

	todolistLenStoreAfter := len(todolist.store)

	if todolistLenStoreBefore == todolistLenStoreAfter {
		t.Errorf("TestTodoList should has len todolistLenStoreBefore %d, less then todolistLenStoreAfter %d", todolistLenStoreBefore, todolistLenStoreAfter)
	}
}

func TestTodoListRemoveItemPositive(t *testing.T) {
	item := Item{
		id:        "1",
		text:      "",
		isChecked: false,
	}

	todolist := Todolist{
		store: []Item{item},
	}

	removedItem, _ := todolist.removeItem(item)

	fmt.Println("removedItem", removedItem)

	if removedItem.id != item.id {
		t.Errorf("Removed item with id %s was not found", item.id)
	}
}

func TestTodoListRemoveItemNegative(t *testing.T) {
	removeItem := Item{
		id:        "1",
		text:      "",
		isChecked: false,
	}

	todolist := Todolist{
		store: []Item{},
	}

	_, err := todolist.removeItem(removeItem)

	fmt.Println("err", err)

	//fmt.Println("removedItem", removedItem)

	if err == nil {
		t.Errorf("Wrong  logic")
	}
}
