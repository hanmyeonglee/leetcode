//go:build ignore

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type (
	Input  = []int
	Output = []string
)

func solution(input Input) Output {
	nums := input
	n := len(nums)
	var ret []string
	l, r := 0, 0
	for l < n {
		start := nums[l]
		for r < n && nums[r]+l-r == start {
			r++
		}
		if l+1 == r {
			ret = append(ret, strconv.Itoa(start))
		} else {
			ret = append(ret, strconv.Itoa(start)+"->"+strconv.Itoa(nums[r-1]))
		}
		l = r
	}

	return ret
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
		{Input{0, 1, 2, 4, 5, 7}, Output{"0->2", "4->5", "7"}},
		{Input{0, 2, 3, 4, 6, 8, 9}, Output{"0", "2->4", "6", "8->9"}},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
