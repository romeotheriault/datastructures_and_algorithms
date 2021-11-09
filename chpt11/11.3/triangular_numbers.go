package main

import f "fmt"

func main() {
	f.Println(triangular_numbers(8))
}

// 1, 3, 6, 10, 15, 21, 28...
func triangular_numbers(a int) int {
	if a == 1 {
		return 1
	}
	return a + triangular_numbers(a-1)
}
