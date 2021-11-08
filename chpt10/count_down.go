package main

import (
	"fmt"
	"time"
)

// Loop countdown
/*
func main() {
	count := os.Args[1]
	c, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for c >= 0 {
		fmt.Println(c)
		c--
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Take off!!!!!!")
}
*/

// Recursive countdown
func main() {
	countdown(10)
}

func countdown(i int) {
	if i < 0 {
		fmt.Println("Take off!!!!")
		return
	}
	fmt.Println(i)
	time.Sleep(1 * time.Second)
	countdown(i - 1)
}
