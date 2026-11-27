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
	Input  = string
	Output = string
)

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func solution(input Input) Output {
	s := input
	b := []byte(s)
	n := len(s)
	l, r := 0, n-1
	for l < r {
		for l < n && !isVowel(b[l]) {
			l++
		}
		for r > 0 && !isVowel(b[r]) {
			r--
		}

		if l < r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}

	return string(b)
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
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
