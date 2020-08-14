package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	inorderNums := make([]int, 0, 8)
	inorderTraversalInner(root, &inorderNums)
	return inorderNums
}

func inorderTraversalInner(root *TreeNode, inorderNums *[]int) {
	if root == nil {
		return
	}
	inorderTraversalInner(root.Left, inorderNums)
	*inorderNums = append(*inorderNums, root.Val)
	inorderTraversalInner(root.Right, inorderNums)
}

//func main() {
//	root := &TreeNode{
//		Val: 1,
//	}
//	node1 := &TreeNode{
//		Val: 2,
//	}
//	node2 := &TreeNode{
//		Val: 3,
//	}
//	root.Left = node1
//	root.Right = node2
//	fmt.Println(inorderTraversal(root))
//}

//非递归实现
func inorderTraversal2(root *TreeNode) []int {
	if root==nil{
		return nil
	}

	res:=make([]int,0,32)
	stack:=make([]*TreeNode,0,32)

	top:=root

	for top!=nil || len(stack)>0{
		for top!=nil {
			stack=append(stack,top)
			res=append(res,top.Val)
			top=top.Left
		}
		if len(stack)>0{
			top=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			top=top.Right
		}
	}
	return res
}


func main(){
	root:=&TreeNode{Val: 1}
	node2:=&TreeNode{Val: 2}
	node3:=&TreeNode{Val: 3}

	root.Right=node2
	node2.Left=node3

	fmt.Println(inorderTraversal2(root))
}