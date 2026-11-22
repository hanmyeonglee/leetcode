//go:build ignore

package main

import (
	"fmt"
	"math/bits"
	"reflect"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type (
	Input  = int
	Output = bool
)

func solution(input Input) Output {
	n := input
	if n <= 0 {
		return false
	}
	m := uint64(n)
	m >>= bits.TrailingZeros64(m)
	return 1418776204833984375%m == 0
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
		{6, true},
		{1, true},
		{14, false},
		{0, false},
		{-6, false},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
