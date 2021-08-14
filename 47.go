package main

import "sort"

// 全排列II
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	track := []int{}
	vis := make(map[int]bool, n)
	res := make([][]int, 0)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			res = append(res, append([]int{}, track...))
			return
		}
		for i, n := range nums {
			if vis[i] || (i > 0 && !vis[i-1] && n == nums[i-1]) {
				continue
			}
			track = append(track, n)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}
