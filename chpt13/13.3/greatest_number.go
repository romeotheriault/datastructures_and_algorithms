package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{}
	start := time.Now()
	for i := 0; i < 20000000; i++ {
		nums = append(nums, rand.Intn(999999))
	}
	duration := time.Since(start)
	fmt.Println("radmon: ", duration)
	//fmt.Println(nums)

	/*
		start := time.Now()
		n := greatest_num_n2(nums)
		duration := time.Since(start)
		fmt.Println("n: ", n, ": duration: ", duration)
	*/

	start = time.Now()
	n := greatest_num_n(nums)
	duration = time.Since(start)
	fmt.Println("n: ", n, ": duration: ", duration)

	start = time.Now()
	n = greatest_num_nlogn(nums)
	duration = time.Since(start)
	fmt.Println("nlog: ", n, ": duration: ", duration)
}

func greatest_num_n2(nums []int) int {
	var max int
	for _, v := range nums {
		for _, iv := range nums {
			if v > max {
				max = v
			} else if iv > max {
				max = iv
			}
		}
	}
	return max
}

func greatest_num_n(nums []int) int {
	max := -1
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func greatest_num_nlogn(nums []int) int {
	qsort(nums, 0, len(nums)-1)
	return nums[len(nums)-1]
}

func qsort(nums []int, l int, r int) {
	if r-l <= 0 {
		return
	}
	pivot := qpartition(nums, l, r)
	qsort(nums, l, pivot-1)
	qsort(nums, pivot+1, r)
}

func qpartition(nums []int, l int, r int) int {
	pivot := r
	r = r - 1

	for {
		for nums[l] < nums[pivot] {
			l++
		}
		for nums[r] > nums[pivot] && r > 0 {
			r--
		}
		if l >= r {
			break
		}
		nums[l], nums[r] = nums[r], nums[l]
		l++
	}
	nums[l], nums[pivot] = nums[pivot], nums[l]
	return l
}
