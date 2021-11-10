package main

import f "fmt"

func main() {
	m := make(map[int]int)
	//f.Println(golomb_sequence(7, m)) // 4
	//f.Println(golomb_sequence(2, m)) // 2
	//f.Println(golomb_sequence(4, m)) // 3
	//f.Println(golomb_sequence(9, m)) // 5
	//f.Println(golomb_sequence(22, m)) // 8
	f.Println(golomb_sequence(50, m)) // 13

}

func golomb_sequence(n int, m map[int]int) int {
	if n == 1 {
		return 1
	}
	if m[n] == 0 {
		m[n] = 1 + golomb_sequence(n-golomb_sequence(golomb_sequence(n-1, m), m), m)
	}
	return m[n]

}
