package avltree

import (
	"fmt"
	"math"
)

type Node interface {
	GetData() int64
	GetLeft() Node
	GetRight() Node
	SetData(int64)
	SetLeft(Node)
	SetRight(Node)
}

type Tree interface {
	GetRoot() Node
	Find(int64) Node
	FindMin() Node
	FindMax() Node
	Insert(int64)
	Delete(int64)
	PopMin() Node
	PopMax() Node
}

type BinaryNode struct {
	left  Node
	right Node
	data  int64
}

func (bn *BinaryNode) GetData() int64 {
	return bn.data
}

func (bn *BinaryNode) GetLeft() Node {
	return bn.left
}

func (bn *BinaryNode) GetRight() Node {
	return bn.right
}

func (bn *BinaryNode) SetData(data int64) {
	bn.data = data
}

func (bn *BinaryNode) SetLeft(n Node) {
	bn.left = n
}

func (bn *BinaryNode) SetRight(n Node) {
	bn.right = n
}

type BinaryTree struct {
	root Node
}

func (bt *BinaryTree) GetRoot() Node {
	return bt.root
}

func (bt *BinaryTree) Find(data int64) Node {
	if bt.root == nil {
		return nil
	}
	tmp := bt.root
	for {
		if tmp.GetData() == data {
			return tmp
		} else if tmp.GetData() < data {
			if tmp.GetRight() != nil {
				tmp = tmp.GetRight()
			} else {
				return nil
			}
		} else {
			if tmp.GetLeft() != nil {
				tmp = tmp.GetLeft()
			} else {
				return nil
			}
		}
	}
}

func (bt *BinaryTree) FindMin() Node {
	if bt.root == nil {
		return nil
	}
	tmp := bt.root
	for {
		if tmp.GetLeft() != nil {
			tmp = tmp.GetLeft()
		} else {
			return tmp
		}
	}
}

func (bt *BinaryTree) FindMax() Node {
	if bt.root == nil {
		return nil
	}
	tmp := bt.root
	for {
		if tmp.GetRight() != nil {
			tmp = tmp.GetRight()
		} else {
			return tmp
		}
	}
}

func (bt *BinaryTree) Insert(data int64) {
	if bt.root == nil {
		bt.root = &BinaryNode{
			left:  nil,
			right: nil,
			data:  data,
		}
		return
	}
	tmp := bt.root
	for {
		if tmp.GetData() == data {
			return
		} else if tmp.GetData() < data {
			if tmp.GetRight() != nil {
				tmp = tmp.GetRight()
			} else {
				tmp.SetRight(&BinaryNode{
					left:  nil,
					right: nil,
					data:  data,
				})
			}
		} else {
			if tmp.GetLeft() != nil {
				tmp = tmp.GetLeft()
			} else {
				tmp.SetLeft(&BinaryNode{
					left:  nil,
					right: nil,
					data:  data,
				})
			}
		}
	}

}

func (bt *BinaryTree) PopMin() Node {
	if bt.root == nil {
		return nil
	}
	tmp := bt.root
	parent := bt.root
	for {
		if tmp.GetLeft() != nil {
			parent = tmp
			tmp = tmp.GetLeft()
		} else {
			if tmp.GetRight() == nil {
				if tmp == bt.root {
					bt.root = nil
				} else {
					parent.SetLeft(nil)
				}
			} else {
				if tmp == bt.root {
					bt.root = bt.root.GetRight()
				} else {
					parent.SetLeft(tmp.GetRight())
				}
			}
			return tmp
		}
	}
}

func (bt *BinaryTree) PopMax() Node {
	if bt.root == nil {
		return nil
	}
	tmp := bt.root
	parent := bt.root
	for {
		if tmp.GetRight() != nil {
			parent = tmp
			tmp = tmp.GetRight()
		} else {
			if tmp.GetLeft() == nil {
				if tmp == bt.root {
					bt.root = nil
				} else {
					parent.SetRight(nil)
				}
			} else {
				if tmp == bt.root {
					bt.root = bt.root.GetLeft()
				} else {
					parent.SetRight(tmp.GetLeft())
				}
			}
			return tmp
		}
	}
}

// 用来简化delete函数
func (bt *BinaryTree) GetData() int64 {
	return bt.root.GetData()
}

func (bt *BinaryTree) GetLeft() Node {
	return bt.root
}

func (bt *BinaryTree) GetRight() Node {
	return bt.root
}

func (bt *BinaryTree) SetLeft(n Node) {
	bt.root = n
}

func (bt *BinaryTree) SetRight(n Node) {
	bt.root = n
}

func (bt *BinaryTree) SetData(data int64) {
	bt.root.SetData(data)
}

func (bt *BinaryTree) Delete(data int64) {
	if bt.root == nil {
		return
	}

	tmp := bt.root
	parent := bt.root

	parentDelete := func() {
		bt.root = nil
	}
	parent2right := func() {
		bt.root = bt.root.GetRight()
	}
	parent2left := func() {
		bt.root = bt.root.GetLeft()
	}
	replace := func(node Node) {
		tt := &BinaryTree{root: bt.root.GetRight()}
		nn := tt.PopMin()
		// todo:下面这句是否可以省略
		bt.root.SetRight(tt.root)
		bt.root.SetData(nn.GetData())
	}
	for {
		if tmp.GetData() == data {
			if tmp.GetLeft() == nil && tmp.GetRight() == nil {
				parentDelete()
			} else if tmp.GetLeft() != nil && tmp.GetRight() == nil {
				parent2left()
			} else if tmp.GetLeft() == nil && tmp.GetRight() != nil {
				parent2right()
			} else {
				replace(tmp)
			}
			return
		} else if tmp.GetData() < data {
			parent = tmp
			tmp = tmp.GetRight()
			if tmp == nil {
				return
			}
			parentDelete = func() {
				parent.SetRight(nil)
			}
			parent2left = func() {
				parent.SetRight(tmp.GetLeft())
			}
			parent2right = func() {
				parent.SetRight(tmp.GetRight())
			}
			replace = func(node Node) {
				tt := &BinaryTree{root: node}
				nn := tt.PopMin()
				parent.SetRight(tt.root)
				node.SetData(nn.GetData())
			}
		} else {
			parent = tmp
			tmp = tmp.GetLeft()
			if tmp == nil {
				return
			}
			parentDelete = func() {
				parent.SetLeft(nil)
			}
			parent2left = func() {
				parent.SetLeft(tmp.GetLeft())
			}
			parent2right = func() {
				parent.SetLeft(tmp.GetRight())
			}
			replace = func(node Node) {
				tt := &BinaryTree{root: node}
				nn := tt.PopMin()
				parent.SetRight(tt.root)
				node.SetData(nn.GetData())
			}
		}
	}
}

func (bt *BinaryTree) Display() {
	PrintTree(bt)
}

// 函数缺陷很明显，当tree过深过大时，性能不行。打印出来的效果也不好看。
func PrintTree(t Tree) {
	if t.GetRoot() == nil {
		return
	}
	array := Tree2Array(t)
	data := Array2Data(array)
	LinePrint(data, 1)
}

func Tree2Array(t Tree) [][]Node {
	if t.GetRoot() == nil {
		return nil
	}
	var temp [][]Node
	temp = make([][]Node, 0, 2)
	depth := 0
	temp = append(temp, []Node{t.GetRoot()})
	for {
		depth += 1
		var nl = make([]Node, 0, 2)
		length := len(temp[depth-1])
		canBreak := true
		for i := 0; i < length; i++ {
			if temp[depth-1][i] != nil {
				nl = append(nl, temp[depth-1][i].GetLeft())
				nl = append(nl, temp[depth-1][i].GetRight())
				if temp[depth-1][i].GetLeft() != nil || temp[depth-1][i].GetRight() != nil {
					canBreak = false
				}
			} else {
				nl = append(nl, nil)
				nl = append(nl, nil)
			}
		}
		if canBreak {
			break
		} else {
			temp = append(temp, nl)
		}
	}
	return temp
}

func Array2Data(bl [][]Node) [][]int64 {
	h := len(bl)
	length := int(math.Exp2(float64(h))) - 1
	tmp := make([][]int64, 0, h)
	for i := 0; i < h; i++ {
		temp := make([]int64, length, length)
		tmp = append(tmp, temp)
	}

	tmp[0][int(math.Exp2(float64(h-1)))-1] = 1
	for i := 1; i < h; i++ {
		jiange := int(math.Exp2(float64(h - i - 1)))
		for j := 0; j < length; j++ {
			if tmp[i-1][j] == 1 {
				tmp[i][j-jiange] = 1
				tmp[i][j+jiange] = 1
			}
		}
	}

	for i := 0; i < h; i++ {
		count := 0
		for j := 0; j < length; j++ {
			if tmp[i][j] == 1 {
				if bl[i][count] != nil {
					tmp[i][j] = bl[i][count].GetData()
				} else {
					tmp[i][j] = 0
				}
				count += 1
			}
		}
	}

	return tmp
}

func LinePrint(data [][]int64, length int) {
	high := len(data)
	if high < 0 {
		return
	}
	l := len(data[0])

	for i := 0; i < high; i++ {
		for j := 0; j < l; j++ {
			if data[i][j] != 0 {
				fmt.Printf("%-*d", length, data[i][j])
			} else {
				fmt.Printf("%-*s", length, " ")
			}
		}
		fmt.Println()
	}
}
