//go:build ignore

package main

import (
	"fmt"
	"reflect"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type (
	Input  = []int
	Output = int
)

func solution(input Input) Output {
	nums := input
	n := len(nums)
	t := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return t - sum
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
		{Input{3, 0, 1}, 2},
		{Input{0, 1}, 2},
		{Input{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
