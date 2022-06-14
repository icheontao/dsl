// Package linearlist implements the linear list.
package linearlist

import (
	"fmt"
	"strings"
)

type List struct {
	elements []interface{}
	size     int
}

// New
func New(vaules ...interface{}) *List {
	list := &List{}
	if len(vaules) < 1 {
		return list
	}
	list.Append(vaules...)
	return list

}

// Append
func (list *List) Append(values ...interface{}) {
	appendNum := len(values)
	if list.size+appendNum > cap(list.elements) {
		list.expansionCap(appendNum)
	}
	for _, elem := range values {
		list.elements[list.size] = elem
		list.size++
	}
}

// Get
func (list *List) Get(index int) (interface{}, bool) {

	if index < 0 || index >= list.size {
		return nil, false
	}
	return list.elements[index], true
}

// Indexof
func (list *List) Indexof(value interface{}) int {

	if list.size > 0 {
		for index, elem := range list.elements {
			if elem == value {
				return index
			}
		}
	}
	return -1

}

// Insert
func (list *List) Insert(index int, values ...interface{}) {

	if index == list.size {
		list.Append(values...)
		return
	}
	if index >= 0 && index < list.size {
		length := len(values)
		if list.size+length > cap(list.elements) {
			list.expansionCap(length)
		}
		copy(list.elements[index+length:], list.elements[index:list.size])
		copy(list.elements[index:index+length], values)
		list.size += length
	}
}

// Remove
func (list *List) Remove(value interface{}) {

	for index, elem := range list.elements {
		if elem == value {
			copy(list.elements[index:], list.elements[index+1:list.size])
			list.elements[list.size-1] = nil
			list.size--
			list.reducionCap()
			return
		}
	}

}

// Len
func (list *List) Len() int {
	return list.size
}

// Empty
func (list *List) Empty() bool {
	return list.size == 0
}

// Clear
func (list *List) Clear() {
	list.elements = []interface{}{}
	list.size = 0
}

// String
func (list *List) String() string {
	ret := []string{}
	for i := 0; i < list.size; i++ {
		ret = append(ret, fmt.Sprintf("%v", list.elements[i]))
	}

	return fmt.Sprintf("LinearList([%s])", strings.Join(ret, ", "))
}

// expansionCap
func (list *List) expansionCap(appendNum int) {

	newCap := cap(list.elements) + appendNum*2
	newElem := make([]interface{}, newCap)
	copy(newElem, list.elements)
	list.elements = newElem
	//fmt.Printf("ExpansionCap: %v, Size: %v, Cap: %v\n", list.elements, list.size, newCap)
}

// reducionCap
func (list *List) reducionCap() {
	if newCap := cap(list.elements) / 2; list.size <= newCap {
		newElem := make([]interface{}, newCap)
		copy(newElem, list.elements)
		list.elements = newElem
		//fmt.Printf("ReducionCap: %v, Size: %v, Cap: %v\n", list.elements, list.size, newCap)
	}
}
