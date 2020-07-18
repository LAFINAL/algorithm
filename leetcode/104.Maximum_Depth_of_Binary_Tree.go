package leetcode


// 일단 전역변수를 써보자.
var count int = 0

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count = count + 1
	a := maxDepth(root.Left)
	b := maxDepth(root.Right)
	if a > b {
		return +a
	} else{
		return 1+b
	}
}
