package main

import (
	"fmt"
)

type tree struct {
	data  int
	right *tree
	left  *tree
}

type bt struct {
	root *tree
}

func (t *bt) insert(data int) *bt {
	if t.root == nil {
		t.root = &tree{data, nil, nil}
		return t
	}
	t.root.insert(data)

	return t
}

func (tr *tree) insert(data int) {
	if data <= tr.data {
		if tr.left == nil {
			tr.left = &tree{data, nil, nil}
			return
		}
		tr.left.insert(data)
		return
	}
	if tr.right == nil {
		tr.right = &tree{data, nil, nil}
		return
	}
	tr.right.insert(data)

}

func print(t *tree) {
	if t == nil {
		return
	}
	print(t.left)
	fmt.Println(t.data)
	print(t.right)
}

func main() {
	t := bt{}
	t.insert(10).insert(23)
	t.insert(9)
	print(t.root)
}
