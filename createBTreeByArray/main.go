package main

import "container/list"

//how to create a binary tree by given array
//[3,9,20,null,null,15,7]
//     3
//    / \
//   9   20
//       / \
//      15 7

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateBTreeByArray(pIntArray []*int) *TreeNode {
	if len(pIntArray) == 0 || pIntArray[0] == nil {
		return nil
	}

	l := list.New()
	root := TreeNode{Val: *pIntArray[0]}
	l.PushBack(root)

	for i := 0; i < len(pIntArray); i++ {

	}
}
