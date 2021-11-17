package main

import (
	"errors"
	"fmt"
)

type treenode struct {
	data  int
	left  *treenode
	right *treenode
}

func createNode(data int) *treenode {
	return &treenode{data: data}
}

func (bst *treenode) search(val int) (*treenode, error) {
	if bst == nil {
		error := fmt.Sprintf("Not found in tree: %d", val)
		return nil, errors.New(error)
	}
	if bst.data == val {
		return bst, nil
	}
	if val < bst.data {
		return bst.left.search(val)
	} else {
		return bst.right.search(val)
	}
}

func (bst *treenode) insert(val int) error {
	if bst == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}
	if val == bst.data { // You can't have duplicate values in a BST.
		return nil
	}
	if val < bst.data {
		if bst.left == nil {
			bst.left = createNode(val)
			return nil
		}
		bst.left.insert(val)
		return nil
	} else {
		if bst.right == nil {
			bst.right = createNode(val)
			return nil
		}
		bst.right.insert(val)
		return nil
	}
}

func (bst *treenode) lift(nodeToDelete *treenode) *treenode {
	if bst.left != nil {
		bst.left = bst.left.lift(nodeToDelete)
		return bst
	} else {
		nodeToDelete.data = bst.data
		return bst.right
	}
}

func (bst *treenode) delete(val int) *treenode {
	if bst == nil {
		fmt.Printf("Not found in tree: %d\n", val)
		return nil
	} else if val < bst.data {
		bst.left = bst.left.delete(val)
		return bst
	} else if val > bst.data {
		bst.right = bst.right.delete(val)
		return bst
	} else if bst.data == val {
		if bst.left == nil {
			return bst.right
		} else if bst.right == nil {
			return bst.left
		} else {
			bst.right = bst.right.lift(bst)
			return bst
		}
	}
	return nil
}

func (bst *treenode) inOrderPrint() {
	if bst == nil {
		return
	}
	bst.left.inOrderPrint()
	fmt.Print(bst.data, ", ")
	bst.right.inOrderPrint()
}

func (bst *treenode) preOrderPrint() {
	if bst == nil {
		return
	}
	fmt.Print(bst.data, ", ")
	bst.left.preOrderPrint()
	bst.right.preOrderPrint()
}

func (bst *treenode) postOrderPrint() {
	if bst == nil {
		return
	}
	bst.left.postOrderPrint()
	bst.right.postOrderPrint()
	fmt.Print(bst.data, ", ")
}

func main() {
	root := createNode(30)

	root.insert(25)
	root.insert(75)

	node, error := root.search(25)
	if error != nil {
		fmt.Println(error)
	}
	if node != nil {
		fmt.Println("Found node: ", node.data)
	}
	root.insert(55)
	root.insert(65)
	root.insert(45)
	root.insert(35)
	node, error = root.search(65)
	if error != nil {
		fmt.Println(error)
	}
	if node != nil {
		fmt.Println("Found node: ", node.data)
	}

	node, error = root.search(45)
	if error != nil {
		fmt.Println(error)
	}
	if node != nil {
		fmt.Println("Found node: ", node.data)
	}

	node, error = root.search(100)
	if error != nil {
		fmt.Println(error)
	}
	if node != nil {
		fmt.Println("Found node: ", node.data)
	}

	root.delete(69)
	root.inOrderPrint()
	fmt.Print("\n")
	root.preOrderPrint()
	fmt.Print("\n")
	root.postOrderPrint()
	fmt.Print("\n")
}
