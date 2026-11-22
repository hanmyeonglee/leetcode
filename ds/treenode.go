package ds

import (
	"fmt"
	"strings"
)

type TreeNode[T comparable] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTree[T comparable](vals ...any) *TreeNode[T] {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	toNode := func(v any) *TreeNode[T] {
		if v == nil {
			return nil
		}
		val, ok := v.(T)
		if !ok {
			panic(fmt.Sprintf("ds.NewTree: expected %T or nil, got %T", *new(T), v))
		}
		return &TreeNode[T]{Val: val}
	}

	root := toNode(vals[0])
	queue := []*TreeNode[T]{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		curr := queue[0]
		queue = queue[1:]

		if i < len(vals) {
			curr.Left = toNode(vals[i])
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			i++
		}
		if i < len(vals) {
			curr.Right = toNode(vals[i])
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			i++
		}
	}
	return root
}

func (t *TreeNode[T]) String() string {
	if t == nil {
		return "[]"
	}

	var out []string
	queue := []*TreeNode[T]{t}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == nil {
			out = append(out, "null")
			continue
		}
		out = append(out, fmt.Sprint(curr.Val))
		queue = append(queue, curr.Left, curr.Right)
	}

	for len(out) > 0 && out[len(out)-1] == "null" {
		out = out[:len(out)-1]
	}

	return "[" + strings.Join(out, ", ") + "]"
}

func (t *TreeNode[T]) Clone() *TreeNode[T] {
	if t == nil {
		return nil
	}
	return &TreeNode[T]{
		Val:   t.Val,
		Left:  t.Left.Clone(),
		Right: t.Right.Clone(),
	}
}

func (t *TreeNode[T]) Equal(other *TreeNode[T]) bool {
	if t == nil || other == nil {
		return t == other
	}
	return t.Val == other.Val && t.Left.Equal(other.Left) && t.Right.Equal(other.Right)
}
