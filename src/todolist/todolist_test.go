package todolist

import (
	"fmt"
	"testing"
)

func TestTodoListAddItem(t *testing.T) {
	item := Item{
		id:        "id",
		text:      "text",
		isChecked: false,
	}

	todolist := Todolist{store: []Item{item}}

	todolistLenStoreBefore := len(todolist.store)

	todolist.addItem(item)

	todolistLenStoreAfter := len(todolist.store)

	if todolistLenStoreBefore == todolistLenStoreAfter {
		t.Errorf("TestTodoList should has len todolistLenStoreBefore %d, less then todolistLenStoreAfter %d", todolistLenStoreBefore, todolistLenStoreAfter)
	}
}

func TestTodoListRemoveItemNegative(t *testing.T) {
	item := Item{
		id:        "1",
		text:      "",
		isChecked: false,
	}

	todolist := Todolist{
		store: []Item{item},
	}

	removedItem, _ := todolist.removeItem(item)

	fmt.Println("", item.id)

	if removedItem.id != item.id{
		t.Errorf("Removed item with id %s was not found", "item.id")
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

	if removedItem.id != item.id{
		t.Errorf("Removed item with id %s was not found", item.id)
	}
}
