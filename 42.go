package main

// 接雨水
// 暴力破解
func trap1(height []int) int {
	n := len(height)
	var ans int
	for i := 0; i < n-1; i++ {
		var leftMax, rightMax int
		for j := i; j < n; j++ {
			rightMax = maxInt(rightMax, height[j])
		}
		for j := i; j >= 0; j-- {
			leftMax = maxInt(leftMax, height[j])
		}
		ans += minInt(leftMax, rightMax) - height[i]
	}
	return ans
}

// 备忘录优化
func trap2(height []int) int {
	n := len(height)
	leftMax := make([]int, n)  // leftMax[i]表示height[0..i]中到最大值
	rightMax := make([]int, n) // rightMax[i]表示height[i...n-1]中的最大值
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = maxInt(height[i], leftMax[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = maxInt(height[i], rightMax[i+1])
	}
	var ans int
	for i := 0; i < n; i++ {
		ans += minInt(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

// 双指针解法
func trap(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	var ans int
	leftMax, rightMax := height[0], height[n-1]
	left, right := 0, n-1
	for left <= right {
		leftMax = maxInt(leftMax, height[left])
		rightMax = maxInt(rightMax, height[right])

		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}
