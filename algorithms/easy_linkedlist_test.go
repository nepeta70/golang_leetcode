package algorithms

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {

	head := &ListNode{Val: 4, Next: &ListNode{
		Val: 5, Next: &ListNode{
			Val: 1, Next: &ListNode{
				Val: 9, Next: nil,
			},
		},
	}}
	expected := &ListNode{Val: 4, Next: &ListNode{
		Val: 1, Next: &ListNode{
			Val: 9, Next: nil,
		},
	}}

	deleteNode(head.Next)

	if !reflect.DeepEqual(head, expected) {
		t.Errorf("Expected %v, got %v", expected, head)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {

	head := &ListNode{Val: 1, Next: &ListNode{
		Val: 2, Next: &ListNode{
			Val: 3, Next: &ListNode{
				Val: 4, Next: &ListNode{
					Val: 5, Next: nil,
				},
			},
		},
	}}
	expected := &ListNode{Val: 1, Next: &ListNode{
		Val: 2, Next: &ListNode{
			Val: 3, Next: &ListNode{
				Val: 5, Next: nil,
			},
		},
	}}
	print(expected.String())
	result := removeNthFromEnd(head, 2)
	print(result.String())
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
func TestRemoveNthFromEnd2(t *testing.T) {

	head := &ListNode{Val: 1, Next: nil}

	result := removeNthFromEnd(head, 1)

	if result != nil {
		t.Errorf("Expected %v, got %v", nil, result)
	}
}

func TestRemoveNthFromEnd3(t *testing.T) {

	head := &ListNode{Val: 1, Next: &ListNode{
		Val: 2, Next: nil}}
	expected := &ListNode{Val: 1, Next: nil}
	print(expected.String())
	result := removeNthFromEnd(head, 1)
	print(result.String())
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
