package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	n := len(nums)
	if n == 0 {
		return 0.0
	}
	if n == 1 {
		return float64(nums[0])
	}
	if n%2 == 0 {
		return float64((nums[n/2-1] + nums[n/2])) / 2
	}
	return float64(nums[(n-1)/2])
}
