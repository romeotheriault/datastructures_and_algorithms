package main

import "fmt"

type node struct {
	data string
	next *node
	prev *node
}

type linkedList struct {
	name string
	head *node
	tail *node
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
	currTail := ll.tail

	if ll.tail == nil {
		ll.head = n
		ll.tail = n
	} else {
		currTail.next = n
		n.prev = currTail
		ll.tail = n
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

func (ll *linkedList) showAllNodesReverse() {
	if ll.tail == nil {
		return
	} else {
		currNode := ll.tail
		for currNode.prev != nil {
			fmt.Print(currNode.data, "--> ")
			currNode = currNode.prev
		}
		fmt.Print(currNode.data, "\n")
	}
}

func (ll *linkedList) getLastElement() string {
	return ll.tail.data
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
		currNode.prev = n
	} else {
		for currNode != nil && curr_index < index-1 {
			currNode = currNode.next
			curr_index++
		}
		//fmt.Println("index: ", index, " curr_index: ", curr_index)
		if currNode != nil {
			n.next = currNode.next
			currNode.next = n
			currNode.next.prev = currNode
			if ll.tail == currNode {
				ll.tail = n
			} else {
				ll.tail = n.next
				n.next.prev = n
			}
		}
	}
}

func (ll *linkedList) delete_at_index(index int) {
	curr_index := 0
	currNode := ll.head

	if index == 0 {
		if currNode.next == nil {
			ll.head = nil
			ll.tail = nil
		} else {
			currNode.next.prev = nil
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
			currNode.next.prev = currNode
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
	fmt.Println(ll.index_of("joey"))
	fmt.Println(ll.index_of("hobbo"))
	fmt.Println(ll.index_of("dog"))
	ll.insert_at_index(0, "maya")
	ll.showAllNodes()
	ll.showAllNodesReverse()
	ll.insert_at_index(0, "romeo")
	ll.showAllNodes()
	ll.showAllNodesReverse()
	fmt.Println(ll.getLastElement())
	ll.insert_at_index(2, "Zahava")
	ll.insert_at_index(5, "Noemi")
	ll.insert_at_index(7, "Gabe")
	ll.insert_at_index(8, "Pocho")
	ll.showAllNodes()
	fmt.Println(ll.getLastElement())
	ll.insert_at_index(10, "Liat")
	ll.showAllNodes()
	ll.showAllNodesReverse()
	/*
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
	*/
}
