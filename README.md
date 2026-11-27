# LeetCode Solutions

Rule #1: I have to beat 75% or more in Runtime
Rule #2: I try to write golang-ideal code
Rule #3: Memory? fuckit

## Setup

Requirements:
- Go (1.21+)
- GNU Make

```bash
git clone git@github.com:hanmyeonglee/leetcode.git
cd leetcode
```

## Usage

### 1. Create a new problem file from `sample.go`

```bash
make prepare https://leetcode.com/problems/two-sum
```

### 2. Format, lint, and run

Runs `gofmt -s -w`, `go vet`, and `go run` in sequence.
If `<filename>` is omitted and there is exactly one new (untracked) `.go` file, it is automatically selected.

```bash
make run two-sum.go
# or auto-detect single new .go file
make run
```

### 3. Format, lint, commit, and push

Commits only the target file with `solve https://leetcode.com/problems/<name>` and pushes to `origin`.
If `<filename>` is omitted and there is exactly one new (untracked) `.go` file, it is automatically selected.

```bash
make git two-sum.go
# or auto-detect single new .go file
make git
```

## Shared Data Structures (`ds/`)

Common LeetCode data structures live under `ds/`. Use a type alias in problem files so `solution` can be copied directly into LeetCode.

```go
import "leetcode/ds"

type ListNode = ds.ListNode[int]
type TreeNode = ds.TreeNode[int]
```

**ListNode**
- `ds.NewList(vals ...T)` — create a linked list
- `(*ListNode).String()`
- `(*ListNode).Clone()`
- `(*ListNode).Equal(other)`

**TreeNode**
- `ds.NewTree[T](vals ...any)` — create a binary tree from LeetCode level-order format (e.g. `ds.NewTree[int](1, 2, 3, nil, 5)`)
- `(*TreeNode).String()`
- `(*TreeNode).Clone()`
- `(*TreeNode).Equal(other)`
