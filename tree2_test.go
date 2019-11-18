package avltree

import (
	"fmt"
	"testing"
)

func TestTree2_PopMin(t *testing.T) {
	tr := new(Tree2)
	tr.Insert(10)
	tr.Insert(9)
	tr.Insert(8)
	tr.Display()

	tr.PopMin()
	tr.Display()
	tr.PopMin()
	tr.Display()
	tr.PopMin()
	fmt.Println("______")
	tr.Display()
}
