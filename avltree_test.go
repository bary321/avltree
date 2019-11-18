package avltree

import (
	"fmt"
	"testing"
)

func TestAVLTree_Find(t *testing.T) {
	tr := new(AVLTree)
	n := new(AVLNode)
	n.Del = false
	n.Depth = 0
	n.Parent = nil
	n.Data = 10
	tr.root = n
	fmt.Println(tr.Find(10))
}
