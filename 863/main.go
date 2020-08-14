package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []int
var k int

func distanceK(root *TreeNode, target *TreeNode, K int) []int {

}

func findChildrenDistanceInK(node *TreeNode, K int, k int) {
	if node == nil {
		return
	}
	if k == K {
		result = append(result, node.Val)
		return
	}
	if k > k {
		return
	}

	findChildrenDistanceInK(node.Left, K, k+1)
	findChildrenDistanceInK(node.Right, K, k+1)
}

func findDistanceKFromRoot(node *TreeNode, target *TreeNode, K int) (bool, int) {
	if node == nil {
		return false, 0
	}
	if node == target {
		return true, 0
	}
	lb, ldistance := findDistanceKFromRoot(node.Left, target, K)
	if lb && ldistance == K {
		result = append(result, node.Val)
		return false, 0
	}

	var rb bool
	var rdistance int
	if lb {
		rb, rdistance = findDistanceKFromRoot(node.Right, target, K-ldistance)
	}

	if rb && rdistance == K {
		result = append(result, node.Val)
	}
	if lb {
		return true, ldistance + 1
	}
	if rb {
		return true, rdistance
	}
	return false, 0
}
