package main

import f "fmt"

func main() {
	f.Println(unique_paths(10, 3))
}

func unique_paths(rows int, cols int) int {
	if rows == 1 || cols == 1 {
		return 1
	}
	return unique_paths(rows-1, cols) + unique_paths(rows, cols-1)
}
