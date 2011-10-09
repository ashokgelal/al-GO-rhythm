package list

import 
	(
		"os"
//		"fmt"
//	if x := f.prev.next; x==nil{
//		fmt.Print("list.head.prev is null")
//	}
//	if x := list.head.prev; x==nil{
//		fmt.Print("list.head.prev is null")
//	}
	)	

func New() *LinkedList {
	list := new(LinkedList)
	list.head = &Node{Value: nil, list: list}
	list.head.next = list.head
	list.head.prev = list.head
	list.length = 0
	return list
}

func (list *LinkedList) Append(val ComparableNodeValue) *Node{
	node := &Node{prev: list.head.prev, next: list.head.next, Value: val, list: list }
	list.head.prev = node
	node.prev.next = node
	list.length++
	return node
}

func (list *LinkedList) Insert(val ComparableNodeValue, index int) (node *Node, err os.Error) {
	if index < 0 || index > list.length {
		return nil, os.NewError("Index out of bounds.")
	}

	current := list.head

	for i := 0; i < index; i++ {
		current = current.next
	}
	
	temp := &Node{prev: current, next: current.next, Value: val, list: list }
	temp.next.prev = temp
	current.next = temp
	list.length++
	return temp, nil
}

func (list *LinkedList) Clear() {
	if !list.IsEmpty() {
		list.head.next = list.head
		list.head.prev = list.head
		list.length = 0
	}
}

func (list *LinkedList) IsEmpty() bool {
	return list.length == 0
}

func (list *LinkedList) Get(index int) (val ComparableNodeValue, err os.Error) {
	if index < 0 || index > list.length {
		err = os.NewError("Index out of bounds.")
	}

	if list.IsEmpty() {
		err = os.NewError("List is empty.")
	}

	if err != nil {
		return
	}

	current := list.head.next

	for i := 0; i < index; i++ {
		current = current.next
	}
	
	val = current.Value

	return
}

func (list *LinkedList) Remove(index int) (err os.Error) {
	if index < 0 || index > list.length {
		err = os.NewError("Index out of bounds.")
	}

	if list.IsEmpty() {
		err = os.NewError("List is empty.")
	}

	if err != nil {
		return
	}

	current := list.head.next

	for i := 0; i < index; i++ {
		current = current.next
	}

	current.prev.next = current.next
	current.next.prev = current.prev
	list.length--
	return nil
}

func (list *LinkedList) RemoveElement(val ComparableNodeValue) bool {
	current := list.head

	for i := 0; i < list.length; i++ {
		current = current.next
		if current.Value.Equals(val) {
			current.prev.next = current.next
			current.next.prev = current.prev
			list.length--
			return true
		}
	}
	return false
}

func (list *LinkedList) GetFirst() (val ComparableNodeValue, err os.Error) {
	if list.IsEmpty() {
		err = os.NewError("List is empty.")
		return
	}

	return list.head.next.Value, nil
}

func (list *LinkedList) PollFirst() (val ComparableNodeValue, err os.Error) {
	val, e := list.GetFirst()
	if e != nil {
		err = e
		return
	}
	list.Remove(0)
	return
}

func (list *LinkedList) Exists(val ComparableNodeValue) bool {
	current := list.head

	for i := 0; i < list.length; i++ {
		current = current.next
		if current.Value.Equals(&val) {
			return true
		}
	}
	return false
}

func (list *LinkedList) Len() int { 
	return list.length
}


type ComparableNodeValue interface{
	Equals(Any) bool
}

type Node struct {
	prev, next *Node
	list       *LinkedList
	Value		ComparableNodeValue
}

type LinkedList struct {
	head   *Node
	length int
}

type Any interface { }
