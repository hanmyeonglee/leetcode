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
	Input  = int
	Output = []int
)

/*
0 = 0

1 = 1

2 = 1
3 = 2

4 = 1
5 = 2
6 = 2
7 = 3

8 = 1
9 = 2
10 = 2
11 = 3
12 = 2
13 = 3
14 = 3
15 = 4

16 = 1
17 = 2
18 = 2
19 = 3
20 = 2
21 = 3
22 = 3
23 = 4
24 = 2
25 = 3
26 = 3
27 = 4
28 = 3
29 = 4
30 = 4
31 = 5

32 = 1
33 = 2
34 = 2
35 = 3
36 = 2
37 = 3
38 = 3
39 = 4
40 = 2
41 = 3
42 = 3
43 = 4
44 = 3
45 = 4
46 = 4
47 = 5
48 = 2
49 = 3
50 = 3
51 = 4
52 = 3
53 = 4
54 = 4
55 = 5
56 = 3
57 = 4
58 = 4
59 = 5
60 = 4
61 = 5
62 = 5
63 = 6

64 = 1
*/

func solution(input Input) Output {
	n := input
	ret := make([]int, n+1)
	if n == 0 {
		return ret
	}

	ret[1] = 1
	b := 1
	for i := 2; i <= n; i++ {
		if i == b<<2 {
			b <<= 1
		}

		v := ret[i-b]
		if i < 3*b {
			ret[i] = v
		} else {
			ret[i] = v + 1
		}
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
		{2, Output{0, 1, 1}},
		{5, Output{0, 1, 1, 2, 1, 2}},
		{8, Output{0, 1, 1, 2, 1, 2, 2, 3, 1}},
	}

	for i, t := range tests {
		test(i, t.input, t.want)
	}
}
