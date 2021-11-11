package main

import f "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10}
	missing := find_missing_num(nums)
	f.Println(missing)
}

func find_missing_num(nums []int) int {
	qsort(nums, 0, len(nums)-1)
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return -1
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
