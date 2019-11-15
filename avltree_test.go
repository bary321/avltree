package avltree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestBinaryTree_PopMax(t *testing.T) {
	r := new(BinaryTree)

	r.Insert(29)
	r.PopMax()

	if r.Find(29) != nil {
		t.Error("Failed")
	}
}

func TestBinaryTree_Insert(t *testing.T) {
	r := new(BinaryTree)

	r.Insert(1)
	r.Insert(23)
	// r.Insert(25)
	for i := 0; i < 5; i++ {
		r.Insert(rand.Int63n(50))
	}

	min := r.FindMin().GetData()
	max := r.FindMax().GetData()
	r.Display()
	fmt.Println("---------------")
	r.PopMin()
	if r.Find(min) != nil {
		t.Error("PopMin Failed")
	}
	r.Display()
	fmt.Println("--------------")
	r.PopMax()
	if r.Find(max) != nil {
		t.Error("PopMax Failed")
	}
	r.Display()
	fmt.Println("--------------")
	r.Delete(23)
	if r.Find(23) != nil {
		t.Error("Delete Failed")
	}
	r.Display()
}

func TestBinaryTree_Insert2(t *testing.T) {
	r := new(BinaryTree)
	r.Insert(30)
	r.Insert(20)
	r.Insert(49)
	r.Insert(59)
	r.Insert(10)
	r.Insert(10)
	r.Insert(10)
	r.Display()
	r.Delete(20)
	r.Display()
}

func TestBinaryTree_Delete(t *testing.T) {
	r := new(BinaryTree)

	r.Insert(29)
	r.Insert(38)
	r.Insert(2)
	r.Delete(29)
	r.Display()
}

func TestBinaryTree_Delete2(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := new(BinaryTree)

		var d int64 = 0
		for i := 0; i < 5000; i++ {
			tmp := rand.Int63n(2000)
			r.Insert(tmp)
			if rand.Intn(30) < 1 {
				d = tmp
			}
		}
		//fmt.Println(d)
		//r.Display()
		r.Delete(d)
		if r.Find(d) != nil {
			fmt.Println(d)
			r.Display()
			t.Error("delete err")
			return
		}
		//fmt.Println("++++++++++++++")
		//r.Display()
		//fmt.Println("______________")
	}
}

func TestBinaryTree_Delete3(t *testing.T) {
	r := new(BinaryTree)
	r.Insert(11)
	r.Insert(8)
	r.Insert(12)
	r.Insert(5)
	r.Display()
	r.Delete(11)
	r.Display()
}

func TestBinaryTree_Delete4(t *testing.T) {
	r := new(BinaryTree)
	r.Insert(5)
	r.Insert(14)
	r.Insert(10)
	r.Insert(19)
	r.Insert(9)
	r.Display()
	r.Delete(14)
	r.Display()
}

func TestLinePrint(t *testing.T) {
	data := [][]int64{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
	}
	LinePrint(data, 1)
}
