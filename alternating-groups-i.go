//go:build ignore

package main

import (
	"fmt"
)

type (
	Input  = []int
	Output = int
)

func solution(input Input) Output {
	colors := input
	n := len(colors)
	cnt := 0
	for i := 0; i < n; i++ {
		a, b, c := colors[i], colors[(i+1)%n], colors[(i+2)%n]
		v := a<<2 | b<<1 | c
		if v == 2 || v == 5 {
			cnt++
		}
		if v == 0 || v == 3 || v == 4 || v == 7 {
			i++
		}
	}

	return cnt
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
		{Input{1, 1, 1}, 0},
		{Input{0, 1, 0, 0, 1}, 3},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
