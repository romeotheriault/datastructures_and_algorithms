package main

import f "fmt"

func main() {
	strings := []string{"one", "two", "three", "four", "five", "six"}
	f.Println(sum_of_strings(strings))
}

func sum_of_strings(s []string) int {
	l := len(s)
	if l == 1 {
		return len(s[0])
	}
	return len(s[0]) + sum_of_strings(s[1:])
}
