//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	lastIdx := len(nums) - 1
	slices.Sort(nums)

	var ret [][]int
	for i := 0; i <= lastIdx-2; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		if a+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if a+nums[lastIdx-1]+nums[lastIdx] < 0 {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			b, c := nums[j], nums[k]
			t := a + b + c
			if t == 0 {
				ret = append(ret, []int{a, b, c})
				for j < k && nums[j] == b {
					j++
				}
				for j < k && nums[k] == c {
					k--
				}
			} else if t > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return ret
}

func test(input []int, want [][]int) {
	got := threeSum(input)
	println("##### Test case #####")
	println("Input: ", fmt.Sprint(input))
	println("Want: ", fmt.Sprint(want))
	println("Got: ", fmt.Sprint(got), "\n")
}

func main() {
	tests := []struct {
		input []int
		want  [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-100, -70, -60, 110, 120, 130, 160}, [][]int{{-100, -60, 160}, {-70, -60, 130}}},
		{[]int{1, 2, 0, 1, 0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, t := range tests {
		test(t.input, t.want)
	}
}
