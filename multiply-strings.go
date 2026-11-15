//go:build ignore

package main

import (
	"fmt"
)

type (
	Input  = [2]string
	Output = string
)

func solution(input Input) Output {
	num1, num2 := input[0], input[1]
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	l1, l2 := len(num1), len(num2)
	l3 := l1 + l2

	ret := make([]byte, l3)
	for i := l1 - 1; i >= 0; i-- {
		a := num1[i] - 48
		idx := l1 - 1 - i
		carry := byte(0)
		for j := l2 - 1; j >= 0; j, idx = j-1, idx+1 {
			b := num2[j] - 48
			c := a*b + carry
			r := c%10 + ret[idx]
			ret[idx] = r % 10
			carry = c/10 + r/10
		}
		ret[idx] = carry
	}

	for i := 0; i < (l3+1)/2; i++ {
		ret[i], ret[l3-i-1] = ret[l3-i-1]+48, ret[i]+48
	}

	j := 0
	for j < l3 {
		if ret[j] != '0' {
			break
		}
		j++
	}

	if j == l3 {
		j--
	}

	return string(ret[j:])
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
		{[2]string{"2", "3"}, "6"},
		{[2]string{"123", "456"}, "56088"},
		{[2]string{"0", "0"}, "0"},
		{[2]string{"9", "99"}, "891"},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
