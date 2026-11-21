//go:build ignore

package main

import (
	"fmt"
)

type (
	Input  = int
	Output = int
)

func solution(input Input) Output {
	n := input
	a, b := 0, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}

	return a
}

func inputToString(input Input) string {
	return fmt.Sprint(input)
}

func outputToString(output Output) string {
	return fmt.Sprint(output)
}

func test(idx int, input Input, want Output) {
	got := solution(input)
	fmt.Printf("##### Test case %d #####\n", idx)
	fmt.Printf("Input: %s\n", inputToString(input))
	fmt.Printf("Want: %s\n", outputToString(want))
	fmt.Printf("Got: %s\n\n", outputToString(got))
}

func main() {
	tests := []struct {
		input Input
		want  Output
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
