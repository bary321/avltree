package avltree

type AVLNode struct {
	Data   int64
	Parent Node
	Depth  int
	Left   Node
	Right  Node
}

func (an *AVLNode) GetData() int64 {
	return an.Data
}

func (an *AVLNode) GetLeft() Node {
	return an.Left
}

func (an *AVLNode) GetRight() Node {
	return an.Right
}

func (an *AVLNode) SetData(data int64) {
	an.Data = data
}

func (an *AVLNode) SetLeft(node Node) {
	an.Left = node
}

func (an *AVLNode) SetRight(node Node) {
	an.Right = node
}
