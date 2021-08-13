package main

// 用两个栈实现队列
type CQueue struct {
	stackIn, stackOut []int
}

func Constructor() CQueue {
	return CQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	// stackIn后边为栈顶
	this.stackIn = append(this.stackIn, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackOut) == 0 {
		if len(this.stackIn) == 0 {
			return -1
		}
		// 从stackIn导入
		this.Pour()
	}
	head := this.stackOut[0]
	this.stackOut = this.stackOut[1:]
	return head
}

func (this *CQueue) Pour() {
	for _, n := range this.stackIn {
		this.stackOut = append(this.stackOut, n)
	}
	this.stackIn = make([]int, 0)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
