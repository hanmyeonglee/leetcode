//go:build ignore

package main

import (
	"fmt"
	"reflect"
	"slices"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type (
	Input  = []int
	Output = []int
)

func solution(input Input) {
	nums := input
	l := 0
	for _, num := range nums {
		if num != 0 {
			nums[l] = num
			l++
		}
	}
	clear(nums[l:])
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
	got := slices.Clone(input)
	solution(got)
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
		{Input{0, 1, 0, 3, 12}, Output{1, 3, 12, 0, 0}},
		{Input{0}, Output{0}},
		{Input{1, 0}, Output{1, 0}},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
