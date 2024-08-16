package main

import "fmt"

// 第K大
// nlogk
func findKthLargest(nums []int, k int) int {
	//fmt.Printf("recursive %v %v\n", nums, k)
	if len(nums) == 0 {
		return -1
	}
	if k > len(nums) || k <= 0 {
		return -1
	}
	pivotIdx := partition(nums, 0, len(nums)-1)
	if pivotIdx == k-1 {
		return nums[pivotIdx]
	} else if pivotIdx > k-1 {
		return findKthLargest(nums[:pivotIdx], k)
	} else {
		return findKthLargest(nums[pivotIdx+1:], k-pivotIdx-1)
	}
}

func partition(nums []int, left, right int) int {
	pivotNum := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] > pivotNum {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func RunKthNumber() {
	fmt.Println(findKthLargest([]int{1, 2, 3}, 3))
	fmt.Println(findKthLargest([]int{1, 2, 5, 4, 5}, 3))
	fmt.Println(findKthLargest([]int{}, 3))
	fmt.Println(findKthLargest([]int{1, 2, 3}, -1))
	fmt.Println(findKthLargest([]int{1, 2, 3}, 4))
	fmt.Println(findKthLargest([]int{1, 2, 3}, 0))
}

/*
func findKthLargest(nums []int, k int) int {
    n := len(nums)
    return quickselect(nums, 0, n - 1, n - k)
}

func quickselect(nums []int, l, r, k int) int{
    if (l == r){
        return nums[k]
    }
    partition := nums[l]
    i := l - 1
    j := r + 1
    for (i < j) {
        for i++;nums[i]<partition;i++{}
        for j--;nums[j]>partition;j--{}
        if (i < j) {
            nums[i],nums[j]=nums[j],nums[i]
        }
    }
    if (k <= j){
        return quickselect(nums, l, j, k)
    }else{
        return quickselect(nums, j + 1, r, k)
    }
}

*/
