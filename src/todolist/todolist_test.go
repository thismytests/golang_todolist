package todolist

import (
	"fmt"
	"testing"
)

func TestFindItemByIdInStorePositive(t *testing.T) {
	index := "1"

	item := Item{
		id:        index,
		text:      "",
		isChecked: false,
	}

	todolist := Todolist{
		store: []Item{item},
	}

	findedItem, _ := todolist.findItemByIdInStore("1")

	if item.id != findedItem.id {
		t.Errorf("Wrong search")
	}
}

func TestTodoListAddItemPositive(t *testing.T) {
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

	storeLengthBefore := len(todolist.store)
	removedItem, _ := todolist.removeItem(item)
	storeLengthAfter := len(todolist.store)

	fmt.Println("removedItem", removedItem)

	if removedItem.id != item.id {
		t.Errorf("Removed item with id %s was not found", item.id)
	}

	if storeLengthBefore > storeLengthAfter {
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

	if err == nil {
		t.Errorf("Wrong  logic")
	}
}

func TestUpdateItemPositive(t *testing.T) {
	updatedItemId := "1"

	item1 := Item{
		id:        updatedItemId,
		text:      "text1",
		isChecked: false,
	}

	item2 := Item{
		id:        "2",
		text:      "text2",
		isChecked: false,
	}

	updatedItem1 := Item{
		id:        updatedItemId,
		text:      "newText",
		isChecked: true,
	}

	todolist := Todolist{
		store: []Item{item1, item2},
	}

	fmt.Println("todolist.store before", todolist.store)
	item, err := todolist.updateItem(updatedItem1.id, updatedItem1)
	fmt.Println("todolist.store after", todolist.store)

	if item != updatedItem1 || err != nil {
		t.Errorf("Update was rejected %s", item.id)
	}
}

func TestUpdateItemNegative(t *testing.T) {
	updatedItemId := "xxx"

	item1 := Item{
		id:        "1",
		text:      "text1",
		isChecked: false,
	}

	item2 := Item{
		id:        "2",
		text:      "text2",
		isChecked: false,
	}

	updatedItem1 := Item{
		id:        updatedItemId,
		text:      "newText",
		isChecked: true,
	}

	todolist := Todolist{
		store: []Item{item1, item2},
	}

	//fmt.Println("todolist.store before", todolist.store)
	_, err := todolist.updateItem(updatedItem1.id, updatedItem1)

	if err == nil {
		t.Errorf("Item with id %s was founded but not exist", updatedItemId)
	}
}
