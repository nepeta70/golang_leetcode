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

	if n == 1 && head.Next == nil {
		return nil
	}

	current := head
	s := []*ListNode{current}
	for {
		if current == nil {
			break
		}
		current = current.Next
		if current == nil {
			break
		}
		s = append(s, current)
	}

	l := len(s)
	indexToDelete := l - n
	if indexToDelete < 0 {
		return head
	}
	if indexToDelete == 0 {
		return head.Next
	}
	toDelete := s[indexToDelete]
	if indexToDelete == l-1 {
		s[l-2].Next = nil
		return head
	}

	toDelete.Val = toDelete.Next.Val
	toDelete.Next = toDelete.Next.Next

	return head
}
