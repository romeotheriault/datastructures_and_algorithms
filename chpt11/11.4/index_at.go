package main

import f "fmt"

func main() {
	s := "abcdefghijklmnopqrstuvwxyz"
	f.Println(index_at(s))
}

func index_at(s string) int {
	if s[0] == 'x' {
		return 0
	}
	return index_at(s[1:]) + 1
}
