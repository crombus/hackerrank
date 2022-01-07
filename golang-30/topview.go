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