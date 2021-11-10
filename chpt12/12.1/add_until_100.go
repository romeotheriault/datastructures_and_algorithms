package main

import f "fmt"

func main() {
	a := []int{2, 3, 4, 5, 20, 21, 19, 5, 6, 7, 13}
	f.Println(add_until_100(a))
}

func add_until_100(a []int) int {
	if len(a) == 0 {
		return 0
	}
	b := add_until_100(a[1:])
	if a[0]+b > 100 {
		return b
	}
	//f.Println(a[0])
	return a[0] + b
}
