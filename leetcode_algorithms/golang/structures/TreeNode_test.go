package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	// 同一个 TreeNode 的不同表达方式
	//            1
	//      	/  \
	//         2    3
	//        / \  /  \
	//       4  5  6   7
	levelOrder = []int{1, 2, 3, 4, 5, 6, 7}
	preOrder   = []int{1, 2, 4, 5, 3, 6, 7} //前序:根左右
	inOrder    = []int{4, 2, 5, 1, 6, 3, 7} //中序:左根右
	postOrder  = []int{4, 5, 2, 6, 7, 3, 1} //左右根
)

func Test_Ints2Tree(t *testing.T) {
	ast := assert.New(t)
	root := PreIn2Tree(preOrder, inOrder) //前序后续构建树

	rootNew := Ints2TreeNode(levelOrder) //层次构建树
	ast.Equal(root, rootNew)
	printTree(root)
}

func Test_Tree2Ints(t *testing.T) {
	ast := assert.New(t)
	root := PreIn2Tree(preOrder, inOrder) //构建树

	ast.Equal(preOrder, Tree2PreOrder(root)) //验证前序遍历切片

	ast.Equal(inOrder, Tree2InOrder(root)) //验证中序遍历切片

	ast.Equal(postOrder, Tree2PostOrder(root)) //验证后序遍历切片

	ast.Equal(postOrder, Tree2PostOrder(root)) //验证后序遍历切片

}

//参考: https://github.com/halfrost/LeetCode-Go/blob/master/structures/TreeNode.go
