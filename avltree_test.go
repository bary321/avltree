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

func TestBinaryTree_Insert(t *testing.T) {
	r := new(BinaryTree)

	for i := 0; i < 5; i++ {
		r.Insert(rand.Int63n(40))
	}

	fmt.Println(r.FindMin())
	fmt.Println(r.FIndMax())
	data := Tree2Array(r)
	if data != nil {
		fmt.Println(data)
	}
	datalist := Array2Data(data)

	fmt.Println(datalist)
}

func TestLinePrint(t *testing.T) {
	data := [][]int64{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
	}
	LinePrint(data)
}
