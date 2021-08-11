package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, n := range nums {
		if idx, e := seen[target-n]; e {
			return []int{idx, i}
		}
		seen[n] = i
	}
	return nil
}
