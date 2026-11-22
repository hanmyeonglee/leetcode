//go:build ignore

package main

import (
	"fmt"
	"reflect"
	"strings"
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
	pattern, s := input[0], input[1]
	sarr := strings.Split(s, " ")
	if len(pattern) != len(sarr) {
		return false
	}

	m := [26]string{}
	for i := 0; i < len(pattern); i++ {
		ch := pattern[i] - 0x61
		w := sarr[i]
		word := m[ch]
		if word == "" {
			for _, existWord := range m {
				if existWord == w {
					return false
				}
			}
			m[ch] = w
		} else {
			if w != word {
				return false
			}
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
		{Input{"abba", "dog cat cat dog"}, true},
		{Input{"abba", "dog cat cat dog"}, true},
		{Input{"aaaa", "dog cat cat fish"}, false},
		{Input{"abba", "dog dog dog dog"}, false},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
