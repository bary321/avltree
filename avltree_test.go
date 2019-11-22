package avltree

import (
	"fmt"
	"math/rand"
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

func TestUpdateDepth(t *testing.T) {
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
	UpdateDepth(root, 1)
	if n1.Depth != 2 {
		t.Error("not 2")
	}
	if n2.Depth != 3 {
		t.Error("not 3")
	}
	if n3.Depth != 2 {
		t.Error("not 2")
	}
}

func TestAVLTree_InsertLeaf(t *testing.T) {
	tr := new(AVLTree)
	tr.InsertLeaf(10)
	tr.InsertLeaf(9)
	tr.InsertLeaf(15)
	tr.InsertLeaf(13)
	tr.InsertLeaf(16)
	tr.Display()
}

func TestAVLTree_InsertLeaf2(t *testing.T) {
	tr := new(AVLTree)
	tr.InsertLeaf(6)
	tr.InsertLeaf(1)
	tr.InsertLeaf(0)
	tr.InsertLeaf(5)
	tr.InsertLeaf(9)
	tr.InsertLeaf(7)
	tr.InsertLeaf(10)
	tr.InsertLeaf(3)
	tr.InsertLeaf(4)
	tr.Display()
}

func TestAVLTree_Insert(t *testing.T) {
	tr := new(AVLTree)
	tr.InsertLeaf(6)
	tr.InsertLeaf(1)
	tr.InsertLeaf(0)
	tr.InsertLeaf(5)
	tr.InsertLeaf(9)
	tr.InsertLeaf(7)
	tr.InsertLeaf(10)
	tr.InsertLeaf(3)
	tr.Display()
	tr.Insert(4)
	tr.Display()
}

func TestAVLTree_Insert2(t *testing.T) {
	tr := new(AVLTree)
	tr.InsertLeaf(6)
	tr.InsertLeaf(1)
	tr.InsertLeaf(0)
	tr.Display()
}

func TestAVLTree_Insert3(t *testing.T) {
	r := new(AVLTree)
	list := make([]int64, 0, 15)
	for i := 0; i < 15; i++ {
		tmp := rand.Int63n(50)
		if tmp == 0 {
			continue
		}
		list = append(list, tmp)
		fmt.Printf("%d, ", tmp)
	}
	fmt.Println()
	for i := 0; i < 15; i++ {
		r.Insert(list[i])
		//fmt.Println("===============")

	}
	//r.Display()

	r.Display()
}

func TestAVLTree_Insert4(t *testing.T) {
	r := new(AVLTree)
	r.Insert(16)
	r.Insert(26)
	r.Insert(11)
	r.Insert(10)
	r.Insert(8)
	r.Insert(7)
	r.Display()
}

func TestMaxDepth(t *testing.T) {
	tr := new(AVLTree)
	tr.InsertLeaf(6)
	tr.InsertLeaf(1)
	tr.InsertLeaf(0)
	tr.InsertLeaf(5)
	tr.InsertLeaf(9)
	tr.InsertLeaf(7)
	tr.InsertLeaf(10)
	tr.InsertLeaf(3)
	tr.Display()
	fmt.Println(MaxDepth(tr.root))
}

func TestAVLTree_Insert5(t *testing.T) {
	r := new(AVLTree)
	r.Insert(5)
	r.Insert(3)
	r.Insert(25)
	r.Insert(2)
	r.Insert(14)
	r.Insert(6)
	r.Insert(25)
	r.Insert(17)
	r.Insert(30)
	r.Insert(17)
	r.Insert(1)
	r.Display()
}

func TestAVLTree_Insert6(t *testing.T) {
	tmp := []int64{3, 11, 41, 36, 14, 18, 43, 19, 38, 12, 8, 46, 32}
	r := new(AVLTree)
	for i := 0; i < 12; i++ {
		//if i == 12 {
		//	r.Display()
		//	fmt.Println()
		//}
		if tmp[i] == 18 {
			fmt.Println()
		}
		fmt.Println(tmp[i])
		r.Insert(tmp[i])
		r.Display()
	}

}
