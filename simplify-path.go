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
	Input  = string
	Output = string
)

func solution(input Input) Output {
	path := input
	n := len(path)
	fp := []string{"/"}
	l, r := 0, 0
	for l < n {
		switch path[l] {
		case '/':
			r++
		default:
			for r < n && path[r] != '/' {
				r++
			}
			chunk := path[l:r]
			switch chunk {
			case ".":
			case "..":
				if len(fp) > 1 {
					fp = fp[:len(fp)-2]
				}
			default:
				fp = append(fp, chunk, "/")
			}
		}
		l = r
	}

	if len(fp) == 1 {
		return "/"
	} else {
		return strings.Join(fp[:len(fp)-1], "")
	}
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
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/../", "/"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
		{"/a/./b/../../c/", "/c"},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
