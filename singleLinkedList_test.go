package singleLinkedList

import "testing"


func TestNewListLen(t *testing.T) {
	l := NewList()
	if l.Len() != 0 {
		t.Fatalf("l.Len() = %d, want %d", l.Len(), 0)
	}
}

func TestListInsert(t *testing.T) {
	l := NewList()
	l.Insert("TEST")
	if l.Len() != 1 {
		t.Fatalf("Insert Failed :l.Len() = %d, want %d", l.Len(), 1)
	}
}