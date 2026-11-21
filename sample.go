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

func inputToString(input Input) string {
	return fmt.Sprint(input)
}

func outputToString(output Output) string {
	return fmt.Sprint(output)
}

func test(idx int, input Input, want Output) {
	got := solution(input)
	fmt.Println("##### Test case ", idx, " #####")
	fmt.Println("Input: ", inputToString(input))
	fmt.Println("Want: ", outputToString(want))
	fmt.Println("Got: ", outputToString(got), "\n")
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
