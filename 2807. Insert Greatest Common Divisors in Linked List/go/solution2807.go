package main

import "fmt"

func main() {
	head := &ListNode{Val: 18}
	head.Next = &ListNode{Val: 6}
	head.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next = &ListNode{Val: 3}
	insertGreatestCommonDivisors(head)
	print(head)
}

func print(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	curr := head
	nxt := head.Next
	for nxt != nil {
		gcdVal := gcd(curr.Val, nxt.Val)
		gcdNode := &ListNode{Val: gcdVal}

		gcdNode.Next = nxt
		curr.Next = gcdNode

		curr = nxt
		nxt = nxt.Next
	}
	return head
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
