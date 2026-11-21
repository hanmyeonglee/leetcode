//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type (
	Input  = [2]*ListNode
	Output = *ListNode
)

func solution(input Input) Output {
	l1, l2 := input[0], input[1]
	root := &ListNode{0, nil}
	curr := root
	carry := 0
	for {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		t := v1 + v2 + carry
		curr.Val = t % 10
		carry = t / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		curr.Next = &ListNode{0, nil}
		curr = curr.Next
	}

	return root
}

func listNodeToString(listNode *ListNode) string {
	curr := listNode
	var arr []string
	for curr != nil {
		arr = append(arr, strconv.Itoa(curr.Val))
		curr = curr.Next
	}

	return strings.Join(arr, " -> ")
}

func inputToString(input Input) string {
	l1, l2 := input[0], input[1]
	return listNodeToString(l1) + " / " + listNodeToString(l2)
}

func outputToString(output Output) string {
	return listNodeToString(output)
}

func test(idx int, input Input, want Output) {
	fmt.Printf("##### Test case %d #####\n", idx)
	fmt.Printf("Input: %s\n", inputToString(input))
	fmt.Printf("Want: %s\n", outputToString(want))
	got := solution(input)
	fmt.Printf("Got: %s\n\n", outputToString(got))
}

func numberToListNode(num int) *ListNode {
	root := &ListNode{num % 10, nil}
	curr := root
	for {
		num /= 10

		if num == 0 {
			break
		}

		curr.Next = &ListNode{num % 10, nil}
		curr = curr.Next
	}

	return root
}

func main() {
	tests := []struct {
		input Input
		want  Output
	}{
		{[2]*ListNode{numberToListNode(342), numberToListNode(465)}, numberToListNode(807)},
		{[2]*ListNode{numberToListNode(0), numberToListNode(0)}, numberToListNode(0)},
		{[2]*ListNode{numberToListNode(9999999), numberToListNode(9999)}, numberToListNode(10009998)},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
