package avltree

import (
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
	if tr.Find(10) == nil {
		t.Error("can't find  10")
	}
}

func TestAVLTree_FindMin(t *testing.T) {
	tr := new(AVLTree)
	root := NewAVLNode(10)
	n1 := NewAVLNode(9)
	n2 := NewAVLNode(8)
	n1.Parent = root
	n2.Parent = n1
	root.SetLeft(n1)
	n1.SetLeft(n2)
	tr.SetRoot(root)
	if tr.FindMin().GetData() != 8 {
		t.Error("not 8")
	}
	n2.Del = true
	if tr.FindMin().GetData() != 9 {
		t.Error("not 9")
	}

	n1.Del = true
	if tr.FindMin().GetData() != 10 {
		t.Error("not 10")
	}

	root.Del = true
	if tr.FindMin() != nil {
		t.Error("not nil")
	}
}

func TestAVLTree_FindMax(t *testing.T) {
	tr := new(AVLTree)
	root := NewAVLNode(10)
	n1 := NewAVLNode(11)
	n2 := NewAVLNode(12)
	n1.Parent = root
	n2.Parent = n1
	root.SetLeft(n1)
	n1.SetLeft(n2)
	tr.SetRoot(root)
	if tr.FindMin().GetData() != 12 {
		t.Error("not 8")
	}
	n2.Del = true
	if tr.FindMin().GetData() != 11 {
		t.Error("not 9")
	}

	n1.Del = true
	if tr.FindMin().GetData() != 10 {
		t.Error("not 10")
	}

	root.Del = true
	if tr.FindMin() != nil {
		t.Error("not nil")
	}
}
