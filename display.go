package avltree

import (
	"fmt"
	"math"
)

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
