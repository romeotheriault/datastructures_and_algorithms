package main

import f "fmt"

func main() {
	j := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 22, 24, 25, 26}
	f.Println(return_even_nums(j))
}

func return_even_nums(a []int) []int {
	if len(a) == 0 {
		return a
	}
	if len(a) == 1 {
		if a[0]%2 == 0 {
			return []int{a[0]}
		} else {
			return []int{}
		}
	}
	if a[0]%2 == 0 {
		return append([]int{a[0]}, return_even_nums(a[1:])...)
	}
	return return_even_nums(a[1:])
}
