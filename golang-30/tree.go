package main

import (
	"container/list"
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

func (t *bt) calllevelorder() {
	h := height(t.root)
	for i := 1; i <= h; i++ {
		printlevelorder(t.root, i)
	}
}

func height(t *tree) int {
	if t == nil {
		return 0
	}
	lh := height(t.left)
	rh := height(t.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func printlevelorder(t *tree, height int) {
	if t == nil {
		return
	}
	if height == 1 {
		fmt.Print(" ", t.data)
	}
	printlevelorder(t.left, height-1)
	printlevelorder(t.right, height-1)
}

func print(t *tree) {
	if t == nil {
		return
	}
	print(t.left)
	fmt.Println(t.data)
	print(t.right)
}

type top struct {
	node *tree
	hd   int
}

func (t *bt) topview() {
	if t.root == nil {
		return
	}
	qu := list.New()
	topview := make(map[int]*tree)
	qu.PushBack(top{t.root, 0})
	topview[qu.Front().Value.(top).hd] = qu.Front().Value.(top).node
	//fmt.Println(sample.Value.(top).hd)
	fmt.Println("top view")
	for qu.Len() > 0 {
		sample := qu.Front()
		qu.Remove(qu.Front())
		for key := range topview {
			if key != sample.Value.(top).hd {
				topview[sample.Value.(top).hd] = sample.Value.(top).node
			}
		}
		if sample.Value.(top).node.left != nil {
			qu.PushBack(top{sample.Value.(top).node.left, sample.Value.(top).hd - 1})
		}
		if sample.Value.(top).node.right != nil {
			qu.PushBack(top{sample.Value.(top).node.right, sample.Value.(top).hd + 1})
		}
	}
	for _, value := range topview {
		fmt.Println(value.data)
	}

}

func main() {
	t := bt{}
	t.insert(10).insert(23)
	t.insert(9).insert(1)
	print(t.root)
	fmt.Println("---------")
	t.calllevelorder()
	t.topview()
}
