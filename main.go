package main

import "fmt"

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	arr := reversePrint(head)
	fmt.Println(arr)
}
