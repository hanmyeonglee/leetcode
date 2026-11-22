//go:build ignore

package main

import (
	"fmt"
	"leetcode/ds"
	"reflect"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type ListNode = ds.ListNode[int]

type (
	Input  = *ListNode
	Output = bool
)

func solution(input Input) Output {
	head := input
	n := 0
	for curr := head; curr != nil; curr = curr.Next {
		n++
	}

	t := (n - 1) / 2
	prev := head
	curr := head.Next
	prev.Next = nil
	for t > 0 {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		t--
	}

	c1, c2 := prev, curr
	if n%2 == 1 {
		c1 = prev.Next
	}

	for c1 != nil && c2 != nil && c1.Val == c2.Val {
		c1 = c1.Next
		c2 = c2.Next
	}

	return c1 == nil && c2 == nil
}

func inputToString(input Input) string {
	return fmt.Sprint(input)
}

func outputToString(output Output) string {
	return fmt.Sprint(output)
}

func compare(want, got Output) bool {
	return reflect.DeepEqual(want, got)
}

func test(idx int, input Input, want Output) {
	got := solution(input.Clone())
	color := colorRed
	if compare(want, got) {
		color = colorGreen
	}
	fmt.Printf("%s##### Test case %d #####\n", color, idx)
	fmt.Printf("Input: %s\n", inputToString(input))
	fmt.Printf("Want: %s\n", outputToString(want))
	fmt.Printf("Got: %s%s\n\n", outputToString(got), colorReset)
}

func main() {
	tests := []struct {
		input Input
		want  Output
	}{
		{ds.NewList(1, 2, 2, 1), true},
		{ds.NewList(1, 2), false},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
