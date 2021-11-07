package main

import f "fmt"

func main() {
	a := []int{0, 2, 4, 6, 8, 10, 7}
	b := []int{6, 7, 8, 9, 10}
	f.Println(intersection(a, b))
}

func intersection(a []int, b []int) []int {
	intersect := []int{}
	m := make(map[int]int)
	for _, v := range a {
		m[v] += 1
	}
	for _, v := range b {
		if m[v] != 0 {
			intersect = append(intersect, v)
		}
	}
	return intersect
}
