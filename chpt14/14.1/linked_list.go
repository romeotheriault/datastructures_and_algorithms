package main

import "fmt"

type node struct {
	data string
	next *node
}

type linkedList struct {
	name string
	head *node
}

func createLinkedList(name string) *linkedList {
	return &linkedList{
		name: name,
	}
}

func (ll *linkedList) addNode(data string) {
	n := &node{
		data: data,
	}
	if ll.head == nil {
		ll.head = n
	} else {
		currNode := ll.head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = n
	}
}

func (ll *linkedList) showAllNodes() {
	if ll.head == nil {
		return
	} else {
		currNode := ll.head
		for currNode.next != nil {
			fmt.Print(currNode.data, "--> ")
			currNode = currNode.next
		}
		fmt.Print(currNode.data, "\n")
	}
}

func (ll *linkedList) getLastElement() string {
	currNode := ll.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	return currNode.data
}

func (ll *linkedList) read(index int) string {
	if ll.head == nil {
		return "No nodes yet!"
	} else {
		curr_index := 0
		currNode := ll.head
		for curr_index < index && currNode != nil {
			currNode = currNode.next
			curr_index++
		}
		if currNode == nil {
			return "There is no node at that index"
		}
		return currNode.data
	}
}

func (ll *linkedList) index_of(value string) int {
	curr_index := 0
	currNode := ll.head
	for currNode != nil {
		if currNode.data == value {
			return curr_index
		}
		currNode = currNode.next
		curr_index++
	}
	return -1
}

func (ll *linkedList) insert_at_index(index int, data string) {
	n := &node{
		data: data,
	}
	curr_index := 0
	currNode := ll.head

	if index == 0 {
		ll.head = n
		n.next = currNode
	} else {
		for currNode != nil && curr_index < index-1 {
			currNode = currNode.next
			curr_index++
		}
		//fmt.Println("index: ", index, " curr_index: ", curr_index)
		if currNode != nil {
			n.next = currNode.next
			currNode.next = n
		}
	}
}

func (ll *linkedList) delete_at_index(index int) {
	curr_index := 0
	currNode := ll.head

	if index == 0 {
		if currNode.next == nil {
			ll.head = nil
		} else {
			ll.head = currNode.next
		}
	} else {
		for currNode != nil && curr_index < index-1 {
			currNode = currNode.next
			curr_index++
		}
		//fmt.Println("index: ", index, " curr_index: ", curr_index)
		if currNode != nil && currNode.next != nil {
			currNode.next = currNode.next.next
		}
	}
}

func main() {
	ll := createLinkedList("hustler")
	ll.addNode("joey")
	ll.addNode("bobby")
	ll.addNode("hobbo")
	fmt.Println(ll.getLastElement())
	ll.showAllNodes()
	fmt.Println(ll.read(0))
	//fmt.Println(ll.index_of("joey"))
	//fmt.Println(ll.index_of("hobbo"))
	//fmt.Println(ll.index_of("dog"))
	ll.insert_at_index(0, "maya")
	ll.insert_at_index(0, "romeo")
	ll.insert_at_index(2, "Zahava")
	ll.insert_at_index(5, "Noemi")
	ll.insert_at_index(7, "Gabe")
	ll.insert_at_index(8, "Pocho")
	fmt.Println(ll.getLastElement())
	ll.insert_at_index(10, "Liat")
	ll.showAllNodes()
	ll.delete_at_index(2)
	ll.showAllNodes()
	ll.delete_at_index(0)
	ll.showAllNodes()
	ll.delete_at_index(6)
	ll.showAllNodes()
	ll.insert_at_index(6, "Liat")
	ll.delete_at_index(7)
	ll.delete_at_index(8)
	ll.showAllNodes()
	fmt.Println(ll.getLastElement())
}
