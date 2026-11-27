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
	Input  = []byte
	Output = []byte
)

func solution(input Input) {
	s := input
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
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
		{Input{'h', 'e', 'l', 'l', 'o'}, Output{'o', 'l', 'l', 'e', 'h'}},
		{Input{'H', 'a', 'n', 'n', 'a', 'h'}, Output{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
