package avltree

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
