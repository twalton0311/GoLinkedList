package singleLinkedList

import "fmt"

type List struct {
	Head   *ListElement
	Tail   *ListElement
	Length *int
}
type ListElement struct {
	Data *interface{}
	Next *ListElement
}

func NewList() List {
	list := List{}
	len := 0
	list.Length = &len
	return list
}

// Return the number of elements in a List
func (list List) Len() int {
	if list.Length == nil {
		*list.Length = 0
	}
	return *list.Length
}

// Traverse the List and return the data
func (list *List) Items() string {
	var output string
	currentElement := list.Head
	output += fmt.Sprintf("[%v", *currentElement.Data)
	// output += elString
	for currentElement.Next != nil {
		currentElement = currentElement.Next
		output += fmt.Sprintf(", %v", *currentElement.Data)
	}
	output += "]"
	return output
}

// Insert an item into a List
func (list *List) Insert(data interface{}) {
	newElement := ListElement{&data, nil}
	if list.Len() == 0 {
		list.Head = &newElement
		list.Tail = &newElement
		*list.Length++
	} else {
		currentElement := list.Head
		for currentElement.Next != nil {
			currentElement = currentElement.Next
		}
		currentElement.Next = &newElement
		list.Tail = &newElement
		*list.Length++
	}
}
