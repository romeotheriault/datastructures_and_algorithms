package main

import "fmt"

type Stack []byte

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(b byte) {
	*s = append(*s, b) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return '0', false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	var stack Stack // create a stack variable of type Stack

	var word string = "banana boat"
	for _, v := range word {
		stack.Push(byte(v))
	}

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Printf("%c", x)
		}
	}
	fmt.Printf("\n")
}
