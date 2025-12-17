package algorithms

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}

	var sb strings.Builder
	for curr := n; curr != nil; curr = curr.Next {
		sb.WriteString(strconv.Itoa(curr.Val))
		if curr.Next != nil {
			sb.WriteString(" -> ")
		}
	}
	sb.WriteString(" -> nil")
	return sb.String()
}

func deleteNode(node *ListNode) {

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	count := 0
	current := head

	for {
		if current == nil {
			break
		}
		current = current.Next
		count++
	}
	toDelete := count - n
	if toDelete < 0 {
		return head
	}
	if count == 1 && toDelete == 1 {
		head = nil
	}
	current = head.Next
	for i := range toDelete {
		if i == toDelete-1 {
			current.Val = current.Next.Val
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
	return head
}
