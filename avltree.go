package avltree

import (
	"fmt"
	"math"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (b *BinaryTree) Find(data int64) *BinaryNode {
	if b.root == nil {
		return nil
	}
	tmp := b.root
	for {
		if tmp.data == data {
			return tmp
		} else if tmp.data < data {
			if tmp.right != nil {
				tmp = tmp.right
			} else {
				return nil
			}
		} else {
			if tmp.left != nil {
				tmp = tmp.left
			} else {
				return nil
			}
		}
	}
}

func (b *BinaryTree) FindMin() *BinaryNode {
	if b.root == nil {
		return nil
	}
	tmp := b.root
	for {
		if tmp.left != nil {
			tmp = tmp.left
		} else {
			return tmp
		}
	}
}

func (b *BinaryTree) FIndMax() *BinaryNode {
	if b.root == nil {
		return nil
	}
	tmp := b.root
	for {
		if tmp.right != nil {
			tmp = tmp.right
		} else {
			return tmp
		}
	}
}

func (b *BinaryTree) Insert(data int64) {
	if b.root == nil {
		b.root = &BinaryNode{
			left:  nil,
			right: nil,
			data:  data,
		}
		return
	}
	tmp := b.root
	for {
		if tmp.data == data {
			return
		} else if tmp.data < data {
			if tmp.right != nil {
				tmp = tmp.right
			} else {
				tmp.right = &BinaryNode{
					left:  nil,
					right: nil,
					data:  data,
				}
			}
		} else {
			if tmp.left != nil {
				tmp = tmp.left
			} else {
				tmp.left = &BinaryNode{
					left:  nil,
					right: nil,
					data:  data,
				}
			}
		}
	}

}

func (b *BinaryTree) Delete(data int64) {

}

func (b *BinaryTree) Display() {

}

func Tree2Array(t *BinaryTree) [][]*BinaryNode {
	if t.root == nil {
		return nil
	}
	var temp [][]*BinaryNode
	temp = make([][]*BinaryNode, 0, 2)
	depth := 0
	temp = append(temp, []*BinaryNode{t.root})
	for {
		depth += 1
		var nl = make([]*BinaryNode, 0, 2)
		length := len(temp[depth-1])
		canBreak := true
		for i := 0; i < length; i++ {
			if temp[depth-1][i] != nil {
				nl = append(nl, temp[depth-1][i].left)
				nl = append(nl, temp[depth-1][i].right)
				if temp[depth-1][i].left != nil || temp[depth-1][i].right != nil {
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

func Array2Data(bl [][]*BinaryNode) [][]int64 {
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
					tmp[i][j] = bl[i][count].data
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
