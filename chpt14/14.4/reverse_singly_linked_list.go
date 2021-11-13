package main

import (
	"fmt"
)

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

func (ll *linkedList) reverse_linked_list() {
	currNode := ll.head
	oldHeadNode := ll.head
	if currNode == nil || currNode.next == nil {
		return
	}

	nextNode := currNode.next
	ll.head = nextNode
	thirdNode := nextNode.next
	nextNode.next = currNode

	for thirdNode != nil {
		currNode = nextNode
		nextNode = thirdNode
		ll.head = nextNode
		thirdNode = nextNode.next
		nextNode.next = currNode
	}
	oldHeadNode.next = nil
}

func main() {
	ll := createLinkedList("hustler")
	ll.addNode("joey")
	ll.addNode("bobby")
	ll.addNode("tappy")
	ll.addNode("flappy")
	ll.addNode("romeo")
	ll.addNode("maya")
	ll.showAllNodes()
	ll.reverse_linked_list()
	ll.showAllNodes()
	ll.reverse_linked_list()
	ll.showAllNodes()
}
