package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "-"
	}

	queue := make([]*TreeNode, 0, 32)
	queue = append(queue, root)

	var str string

	for len(queue) > 0 {
		if isAllEmptyNode(queue) {
			break
		}

		levelNodeCount := len(queue)
		for levelNodeCount > 0 {
			node := queue[0]
			if node == nil {
				queue = append(queue, nil)
				queue = append(queue, nil)
			} else {
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
			if node == nil {
				str = str + "-" + ","
			} else {
				str = str + strconv.Itoa(node.Val) + ","
			}

			queue = queue[1:]
			levelNodeCount--
		}
	}

	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 || data[0] == '-' { //root is nil
		return nil
	}

	//len(data) 一定是大于1的
	if data[len(data)-1] == ',' {
		data = data[:len(data)-1] //去掉最后一","
	}

	//根据","切割
	slice := strings.Split(data, ",")
	newSlice := make([]string, len(slice)+1)
	for i := range slice {
		newSlice[i+1] = slice[i]
	}

	nodeMap := make(map[int]*TreeNode)

	root := &TreeNode{Val: string2int(newSlice[1])}
	nodeMap[1] = root

	for i := 1; i < len(newSlice); i++ {
		node := nodeMap[i]
		if 2*i < len(newSlice) && newSlice[2*i] != "-" {
			node.Left = &TreeNode{Val: string2int(newSlice[2*i])}
			nodeMap[2*i] = node.Left
		}
		if 2*i+1 < len(newSlice) && newSlice[2*i+1] != "-" {
			node.Right = &TreeNode{Val: string2int(newSlice[2*i+1])}
			nodeMap[2*i+1] = node.Right
		}
	}
	return root
}

func isAllEmptyNode(queue []*TreeNode) bool {
	for i := range queue {
		if queue[i] != nil {
			return false
		}
	}
	return true
}

func string2int(str string) int {
	stri, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return stri
}

func main() {
	root := &TreeNode{Val: 1}
	//node2 := &TreeNode{Val: 2}
	//node3 := &TreeNode{Val: 3}

	//root.Right = node3
	//node3.Left = node2

	c := Constructor()
	serial := c.serialize(root)

	newRoot := c.deserialize(serial)

	dfs(newRoot)
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	dfs(node.Left)
	dfs(node.Right)
}
