package avltree

type AVLNode struct {
	Data   int64
	Parent Node
	Depth  int
	Left   Node
	Right  Node
	Del    bool
}

func NewAVLNode(data int64) *AVLNode {
	a := new(AVLNode)
	a.Data = data
	a.Parent = nil
	a.Depth = 0
	a.Left = nil
	a.Right = nil
	a.Del = false
	return a
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

func (an *AVLNode) IsDel() bool {
	return an.Del
}

func (an *AVLNode) Delete() {
	an.Del = true
}

func (an *AVLNode) Recover() {
	an.Del = false
}

func (an *AVLNode) GetParent() Node {
	return an.Parent
}

func (an *AVLNode) SetParent(node Node) {
	an.Parent = node
}

func (an *AVLNode) GetDepth() int {
	return an.Depth
}

func (an *AVLNode) SetDepth(d int) {
	an.Depth = d
}
