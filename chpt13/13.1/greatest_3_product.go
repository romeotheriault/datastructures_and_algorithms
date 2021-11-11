package main

import f "fmt"

func main() {
	nums := []int{3, 20, 33, 1, 0, 10, 99, 0, 1, 2, 7, 8}
	f.Println(nums)
	qsort(nums, 0, len(nums)-1)
	f.Println(nums)
	alen := len(nums) - 1
	f.Println("Greatest product of 3 nums: ", nums[alen-2]*nums[alen-1]*nums[alen])
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
