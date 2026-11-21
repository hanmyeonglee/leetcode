//go:build ignore

package main

import (
	"fmt"
)

type (
	Input  = [][]int
	Output = int
)

func solution(input Input) Output {
	intervals := input
	cnt := 0
	for i, tuple1 := range intervals {
		x1, y1 := tuple1[0], tuple1[1]
		for _, tuple2 := range intervals[i+1:] {
			x2, y2 := tuple2[0], tuple2[1]
			if y1 < x2 || y2 < x1 {
				continue
			}
			cnt++
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
		{Input{{1, 2}, {2, 3}, {3, 4}}, 2},
		{Input{{1, 5}, {2, 4}, {3, 6}}, 3},
		{Input{{1, 2}, {3, 4}, {5, 6}}, 0},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
