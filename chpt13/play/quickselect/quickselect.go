package main

import f "fmt"

func main() {
	nums := []int{55, 66, 77, 3, 4, 9, 2, 0, 33, 22, 7, 8, 11, 5, 10}
	f.Println(nums)
	qsort(nums, 0, len(nums)-1)
	f.Println(nums)
}

func qselect(nums []int, l int, r int) {
	if r-l <= 0 {
		return
	}
	var pivot int = qpartition(nums, l, r)
	//f.Println(pivot)
	qsort(nums, l, pivot-1)
	qsort(nums, pivot+1, r)
}

func qpartition(nums []int, l int, r int) int {
	var pivot int = r
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
		// Swap left and right values
		nums[l], nums[r] = nums[r], nums[l]
		l++
	}

	// Swap the pivot with the left pointer (it's final destination)
	nums[pivot], nums[l] = nums[l], nums[pivot]
	return l
}
