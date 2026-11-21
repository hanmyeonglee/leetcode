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
	Input  = [][]int
	Output = bool
)

func solution(input Input) Output {
	points := input
	p1, p2, p3 := points[0], points[1], points[2]
	x1, y1 := p1[0], p1[1]
	x2, y2 := p2[0], p2[1]
	x3, y3 := p3[0], p3[1]

	return (y2-y1)*(x3-x2) != (y3-y2)*(x2-x1)
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
		{Input{{1, 1}, {2, 3}, {3, 2}}, true},
		{Input{{1, 1}, {2, 2}, {3, 3}}, false},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
