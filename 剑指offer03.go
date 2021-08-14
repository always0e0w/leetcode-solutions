package main

// 数组中重复的数字
func findRepeatNumber(nums []int) int {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return n
		}
		seen[n] = true
	}
	return -1
}
