package main

import (
	"encoding/json"
	"fmt"
)

type BTNode struct {
	Val         int
	Left, Right *BTNode
}

func ParseBinaryTree(s string) *BTNode {
	vals := make([]*int, 0)
	_ = json.Unmarshal([]byte(s), &vals)
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	var queue []*BTNode
	root := &BTNode{Val: *vals[0]}
	queue = append(queue, root)
	i := 1
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if i <= len(vals)-1 && vals[i] != nil {
			n.Left = &BTNode{Val: *vals[i]}
			queue = append(queue, n.Left)
		}
		i++
		if i <= len(vals)-1 && vals[i] != nil {
			n.Right = &BTNode{Val: *vals[i]}
			queue = append(queue, n.Right)
		}
		i++
	}
	return root
}
func (root *BTNode) String() string {
	re := make([]*int, 0)
	var queue []*BTNode
	if root != nil {
		re = append(re, &root.Val)
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		ns := queue[0:]
		queue = nil
		for _, n := range ns {
			if n.Left == nil {
				re = append(re, nil)
			} else {
				re = append(re, &n.Left.Val)
				queue = append(queue, n.Left)
			}
			if n.Right == nil {
				re = append(re, nil)
			} else {
				re = append(re, &n.Right.Val)
				queue = append(queue, n.Right)
			}
		}
	}
	// 去除结尾nil
	i := len(re) - 1
	for i >= 0 && re[i] == nil {
		i--
	}
	re = re[:i+1]
	b, _ := json.Marshal(re)
	return string(b)
}

func PathSum3() {

}
func RunPathSum3() {
	fmt.Println(ParseBinaryTree("[10,5,-3,3,2,null,11,3,-2,null,1]"))
}
