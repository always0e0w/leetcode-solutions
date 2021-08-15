package main

// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	arr := helper(head, []int{})
	return arr
}

func helper(head *ListNode, arr []int) []int {
	if head == nil {
		return arr
	}
	arr = append(helper(head.Next, arr), head.Val)
	return arr
}
