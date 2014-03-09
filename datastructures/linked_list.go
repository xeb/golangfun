package datastructs

import (
// "fmt"
)

type Element struct {
	Prev *Element
	Next *Element
	Item interface{}
}

type LinkedList struct {
	Elements []*Element
}

func NewLinkedList() *LinkedList {
	eles := make([]*Element, 0)
	return &LinkedList{eles}
}

func (ll *LinkedList) Get(index int) (e *Element) {
	e = ll.Elements[index]
	return
}

func (ll *LinkedList) Append(item interface{}) {
	ele := &Element{Item: item}

	ll.Elements = append(ll.Elements, ele)

	length := len(ll.Elements)

	if length > 1 {
		ele.Prev = ll.Elements[length-2]
		if ele.Prev != nil {
			ele.Prev.Next = ele
		}
	}

}
