package main

import (
	f "fmt"
	s "strings"
)

func main() {
	f.Printf("%c\n", missing_letter("abcdefghijklmnoqrstuvwxyz"))
}

func missing_letter(a string) rune {

	m := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}

	a = s.ToLower(a)
	for _, v := range a {
		m[v]++
	}

	for k, v := range m {
		if v == 0 {
			return k
		}
	}
	return '0'
}
