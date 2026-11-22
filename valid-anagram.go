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
	Input  = [2]string
	Output = bool
)

func solution(input Input) Output {
	s, t := input[0], input[1]
	if len(s) != len(t) {
		return false
	}

	var ms [26]uint16
	var mt [26]uint16
	for i := 0; i < len(s); i++ {
		ms[s[i]-0x61] += 1
		mt[t[i]-0x61] += 1
	}

	for i := 0; i < 26; i++ {
		if ms[i] != mt[i] {
			return false
		}
	}

	return true
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
		{Input{"anagram", "nagaram"}, true},
		{Input{"rat", "car"}, false},
		{Input{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba",
		}, false},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
