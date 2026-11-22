package ds

import (
	"fmt"
	"strings"
)

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

func NewList[T comparable](vals ...T) *ListNode[T] {
	dummy := &ListNode[T]{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode[T]{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func (l *ListNode[T]) String() string {
	if l == nil {
		return "nil"
	}

	var b strings.Builder
	visited := make(map[*ListNode[T]]bool)
	for curr := l; curr != nil; curr = curr.Next {
		if b.Len() > 0 {
			b.WriteString(" -> ")
		}
		if visited[curr] {
			b.WriteString("(cycle)")
			break
		}
		visited[curr] = true
		b.WriteString(fmt.Sprint(curr.Val))
	}
	return b.String()
}

func (l *ListNode[T]) Clone() *ListNode[T] {
	dummy := &ListNode[T]{}
	dst := dummy
	for curr := l; curr != nil; curr = curr.Next {
		dst.Next = &ListNode[T]{Val: curr.Val}
		dst = dst.Next
	}
	return dummy.Next
}

func (l *ListNode[T]) Equal(other *ListNode[T]) bool {
	a, b := l, other
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}
