package main

import "fmt"

// O(N) O(N)
// 额外约束: 1到N之间的数字
func FindDuplicateNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i]-1 {
			if nums[nums[i]-1] == nums[i] {
				return nums[i]
			}
			swap(nums, i, nums[i]-1)
		}
	}
	return -1
}

func RunFindDuplicateNumber() {
	fmt.Println(FindDuplicateNumber([]int{1, 2, 3, 3}))
	fmt.Println(FindDuplicateNumber([]int{1, 4, 4, 3}))
	fmt.Println(FindDuplicateNumber([]int{1, 2, 3, 4}))
}
