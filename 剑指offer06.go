package main

// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	arr := make([]int, 0)
	var f func(*ListNode)
	f = func(head *ListNode) {
		if head == nil {
			return
		}
		f(head.Next)
		arr = append(arr, head.Val)
	}
	f(head)
	return arr
}
