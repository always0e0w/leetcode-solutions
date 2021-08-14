package main

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	slow, fast := 0, 1
	for fast < n {
		if nums[fast] > nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
