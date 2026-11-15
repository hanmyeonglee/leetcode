//go:build ignore

package main

import (
	"fmt"
)

type (
	Input  = any
	Output = any
)

func solution(input Input) Output {
	return nil
}

func test(idx int, input Input, want Output) {
	got := solution(input)
	println("##### Test case ", idx, " #####")
	println("Input: ", fmt.Sprint(input))
	println("Want: ", fmt.Sprint(want))
	println("Got: ", fmt.Sprint(got), "\n")
}

func main() {
	tests := []struct {
		input Input
		want  Output
	}{
		// {Input{}, Output{}},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
