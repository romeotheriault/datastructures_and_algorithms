package main

import f "fmt"

func main() {
	f.Println("hi")
	j := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f.Println((j))
}

func return_even_nums(a []int) []int {

}

func return_even_nums(a []int) []int {
	b := []int{}
	if len(a) == 0 {
		return b
	}
	if len(a) == 1 {
		if a[0]%2 == 0 {
			return append(b, a[0])
		} else {
			return b
		}
	}
	if a[0]%2 == 0 {
		return return_even_nums(append(b, a[0]))
	} else {
		return return_even_nums(b)
	}
}
