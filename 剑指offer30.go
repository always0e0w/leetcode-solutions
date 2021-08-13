package main

import "math"

// 包含min函数的栈
type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
/*
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}
*/

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	min := this.Min()
	if x < min {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, min)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return math.MaxInt64
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.minStack) == 0 {
		return math.MaxInt64
	}
	return this.minStack[len(this.minStack)-1]
}
