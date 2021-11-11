package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[string]int)
	fmt.Println(unique_paths(30, 4, m))
}

func unique_paths(rows int, cols int, m map[string]int) int {
	//fmt.Println("recurse")
	if rows == 1 || cols == 1 {
		return 1
	}
	if m[strconv.Itoa(rows)+strconv.Itoa(cols)] == 0 {
		m[strconv.Itoa(rows)+strconv.Itoa(cols)] = unique_paths(rows-1, cols, m) + unique_paths(rows, cols-1, m)
	}
	return m[strconv.Itoa(rows)+strconv.Itoa(cols)]
}
