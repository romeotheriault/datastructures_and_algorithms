package main

import f "fmt"

func main() {
	strings := []string{"a", "b", "c", "d", "e", "a", "f", "g"}
	m := make(map[string]int)

	for _, v := range strings {
		if m[v] == 0 {
			m[v]++
		} else {
			f.Println(v)
		}
	}
}
