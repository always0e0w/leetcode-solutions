package main

// 全排列
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	seen := make(map[int]bool, len(nums))
	var backtrack func([]int, []int)
	backtrack = func(nums []int, track []int) {
		if len(track) == len(nums) {
			path := make([]int, len(track))
			copy(path, track)
			res = append(res, path)
			return
		}
		for _, n := range nums {
			if seen[n] {
				continue
			}
			track = append(track, n)
			seen[n] = true
			backtrack(nums, track)
			seen[track[len(track)-1]] = false
			track = track[:len(track)-1]
		}
	}
	track := make([]int, 0)
	backtrack(nums, track)
	return res
}
