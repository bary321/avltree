package avltree

import (
	"fmt"
	"reflect"
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

func TestAVLTree_Delete(t *testing.T) {
	tr := new(AVLTree)
	root := NewAVLNode(10)
	n1 := NewAVLNode(11)
	n2 := NewAVLNode(12)
	n1.Parent = root
	n2.Parent = n1
	root.SetRight(n1)
	n1.SetRight(n2)
	tr.SetRoot(root)
	if tr.PopMax().GetData() != 12 {
		t.Error("not 12")
	}
	if tr.PopMin().GetData() != 10 {
		t.Error("not10")
	}
	tr.Delete(11)
	if tr.Find(11) != nil {
		t.Error("find err")
	}
}

func TestRightSingleRotation(t *testing.T) {
	tr := new(AVLTree)
	root := NewAVLNode(11)
	n1 := NewAVLNode(9)
	n2 := NewAVLNode(8)
	n1.Parent = root
	n2.Parent = n1
	root.SetLeft(n1)
	n1.SetLeft(n2)
	tr.SetRoot(root)
	n3 := NewAVLNode(13)
	root.Right = n3
	n3.Parent = root
	n4 := NewAVLNode(10)
	n1.SetRight(n4)
	n4.Parent = n1
	tr.Display()
	fmt.Println("-------------")
	tr.root = RightSingleRotation(root)
	tr.Display()
	fmt.Println("-------------")
	tr.root = LeftSingleRotation(tr.root.(*AVLNode))
	tr.Display()
}

func TestRightDoubleRotation(t *testing.T) {
	tr := new(AVLTree)
	root := NewAVLNode(11)
	n1 := NewAVLNode(9)
	n2 := NewAVLNode(8)
	n1.Parent = root
	n2.Parent = n1
	root.SetLeft(n1)
	n1.SetLeft(n2)
	tr.SetRoot(root)
	n3 := NewAVLNode(13)
	root.Right = n3
	n3.Parent = root
	n4 := NewAVLNode(10)
	n1.SetRight(n4)
	n4.Parent = n1
	tr.Display()
	tr.root = RightDoubleRotation(root)
	tr.Display()
}

func TestLeftDoubleRotation(t *testing.T) {
	tr := new(AVLTree)
	root := NewAVLNode(10)
	n1 := NewAVLNode(15)
	n2 := NewAVLNode(13)
	n1.Parent = root
	n2.Parent = n1
	root.SetRight(n1)
	n1.SetLeft(n2)
	tr.SetRoot(root)
	n3 := NewAVLNode(9)
	root.Left = n3
	n3.Parent = root
	n4 := NewAVLNode(16)
	n1.SetRight(n4)
	n4.Parent = n1
	tr.Display()
	tr.root = LeftDoubleRotation(root)
	tr.Display()
}

func TestIsNil(t *testing.T) {
	a := make([]Node, 0, 1)
	a = append(a, nil)
	fmt.Println(IsNil(a[0]))
	fmt.Println(reflect.ValueOf(a[0]))
	fmt.Println(reflect.TypeOf(a[0]))
}
