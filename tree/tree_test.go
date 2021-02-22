package tree

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	root := initTree(&TreeNode{Val: 1})

	preOrderTraversalRec(root)
	fmt.Println(preOrderTraversalNonRec(root))

	inOrderTraversalRec(root)
	fmt.Println(inOrderTraversalNonRec(root))

	postOrderTraversalRec(root)
	fmt.Println(postOrderTraversalNonRec(root))
}
