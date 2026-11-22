//go:build ignore

package main

import (
	"fmt"
	"leetcode/ds"
	"reflect"
	"strconv"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type TreeNode = ds.TreeNode[int]

type (
	Input  = *TreeNode
	Output = []string
)

func traversal(node *TreeNode, parent string) []string {
	var me string
	if parent == "" {
		me = strconv.Itoa(node.Val)
	} else {
		me = parent + "->" + strconv.Itoa(node.Val)
	}
	if node.Left == nil && node.Right == nil {
		return []string{me}
	}

	ret := []string{}
	if node.Left != nil {
		ret = append(ret, traversal(node.Left, me)...)
	}
	if node.Right != nil {
		ret = append(ret, traversal(node.Right, me)...)
	}

	return ret
}

func solution(input Input) Output {
	root := input
	return traversal(root, "")
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
	got := solution(input)
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
		{ds.NewTree[int](1, 2, 3, nil, 5), Output{"1->2->5", "1->3"}},
		{ds.NewTree[int](1), Output{"1"}},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
