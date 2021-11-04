package main

import "fmt"

// ========================{
// Find the greatest number
// ========================
func main() {
	nums := []int{2, 3, 4}
	fmt.Println(greatestNumberQuadratic(nums))
	fmt.Println(greatestNumberLinear(nums))
}

func greatestNumberLinear(a []int) int {
	var greatest int = a[0]
	for _, i := range a {
		if i > greatest {
			greatest = i
		}
	}
	return greatest
}

func greatestNumberQuadratic(a []int) int {
	for _, inum := range a {
		is_inum_greatest := true

		for _, jnum := range a {
			if jnum > inum {
				is_inum_greatest = false
			}
		}
		if is_inum_greatest {
			return inum
		}
	}
	return 0
}
