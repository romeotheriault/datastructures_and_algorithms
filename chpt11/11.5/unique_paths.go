package main

import f "fmt"

func main() {
	f.Println(unique_paths(30, 4))
}

func unique_paths(rows int, cols int) int {
	//f.Println("recurse")
	if rows == 1 || cols == 1 {
		return 1
	}
	return unique_paths(rows-1, cols) + unique_paths(rows, cols-1)
}
